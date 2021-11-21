package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/nsf/termbox-go"
)

func convertTermboxKeyToString(key termbox.Key) string {
	switch key {
	case termbox.KeySpace:
		return "key_space"
	case termbox.KeyBackspace:
		return "key_backspace"
	case termbox.KeyDelete:
		return "key_delete"
	case termbox.KeyInsert:
		return "key_insert"
	case termbox.KeyEnter:
		return "key_enter"
	case termbox.KeyF1:
		return "key_f1"
	case termbox.KeyF2:
		return "key_f2"
	case termbox.KeyF3:
		return "key_f3"
	case termbox.KeyF4:
		return "key_f4"
	case termbox.KeyF5:
		return "key_f5"
	case termbox.KeyF6:
		return "key_f6"
	case termbox.KeyF7:
		return "key_f7"
	case termbox.KeyF8:
		return "key_f8"
	case termbox.KeyF9:
		return "key_f9"
	case termbox.KeyF10:
		return "key_f10"
	case termbox.KeyF11:
		return "key_f11"
	case termbox.KeyF12:
		return "key_f12"
	case termbox.KeyHome:
		return "key_home"
	case termbox.KeyEnd:
		return "key_end"
	case termbox.KeyPgup:
		return "key_page_up"
	case termbox.KeyPgdn:
		return "key_page_down"
	case termbox.KeyArrowUp:
		return "key_arrow_up"
	case termbox.KeyArrowDown:
		return "key_arrow_down"
	case termbox.KeyArrowLeft:
		return "key_arrow_left"
	case termbox.KeyArrowRight:
		return "key_arrow_right"
	default:
		return "key_undefined"
	}
}

func convertTermboxEventToString(e termbox.Event) string {
	if e.Ch == 0 {
		return convertTermboxKeyToString(e.Key)
	}

	return "key_" + string(e.Ch)
}

func findCommand(ctx context.Context, config Config, event termbox.Event) *exec.Cmd {
	for k, v := range config.Commands {
		if k != convertTermboxEventToString(event) {
			continue
		}
		return exec.CommandContext(ctx, v.Name, v.Args...)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.New(os.Stderr, "error: ", 0).Fatal(err)
	}
}

func run() error {
	config, err := loadConfig()

	if err != nil {
		return fmt.Errorf("failed to load config file: %w", err)
	}
	if err := termbox.Init(); err != nil {
		return fmt.Errorf("failed to initialize termbox: %w", err)
	}

	defer termbox.Close()

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

RUNLOOP:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				break RUNLOOP
			}

			wg.Add(1)

			go func() {
				command := findCommand(ctx, config, event)

				if command != nil {
					fmt.Printf("%s\tcommand: %s\n", convertTermboxEventToString(event), strings.Join(command.Args, " "))
					if err := command.Run(); err != nil {
						fmt.Printf("%s\terror: %v\n", convertTermboxEventToString(event), err)
					}
				} else {
					fmt.Printf("%s\tempty\n", convertTermboxEventToString(event))
				}

				wg.Done()
			}()
		}
	}

	cancel()
	wg.Wait()

	return nil
}
