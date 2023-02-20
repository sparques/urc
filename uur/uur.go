// uur is a package for using an HC-12 UART radio connected to a CH340 USB<->UART; the main trick here is the CH340
// has a DTR pin broken out, we have that connected to the HC-12's set pin so by flipping the DTR line we can
// switch between command mode and transpart transmission mode
package uur

import (
	"github.com/pkg/term"
)

const (
	commandModeBaud = 9600
)

type Radio struct {
	*term.Term
	// commandMode is true when we're in command mode
	commandMode bool

	txOptions []func(*term.Term) error
}

func New(path string, opts ...func(*term.Term) error) (*Radio, error) {
	t, err := term.Open(path, opts...)
	if err != nil {
		return nil, err
	}
	return &Radio{
		Term: t,
	}, nil
}

// Mode returns true if we're in command mode, false if in transparent data
// transmission mode
func (r *Radio) Mode() bool {
	d, _ := r.DTR()
	return d
}

func (r *Radio) CommandMode() {
	r.SetDTR(true)
	r.Term.SetOption(term.Speed(commandModeBaud))
	r.commandMode = true
}

func (r *Radio) SetOption(opts ...func(*term.Term) error) {
	r.txOptions = opts
	r.Term.SetOption(opts...)
}

func (r *Radio) TransmitMode() {
	r.SetDTR(false)
	r.commandMode = false
	r.Term.SetOption(r.txOptions...)
}
