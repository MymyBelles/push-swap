package main

import (
	"reflect"
	"testing"
)

func TestArgumentsValides(t *testing.T) {
	pile, valide := RecupererArguments([]string{"3 2 1"})

	if !valide || !reflect.DeepEqual(pile, []int{3, 2, 1}) {
		t.Errorf("Arguments valides mal traités : %v", pile)
	}
}

func TestArgumentsNegatifs(t *testing.T) {
	pile, valide := RecupererArguments([]string{"-3 0 2"})

	if !valide || !reflect.DeepEqual(pile, []int{-3, 0, 2}) {
		t.Errorf("Nombres négatifs mal traités : %v", pile)
	}
}

func TestArgumentsDoublons(t *testing.T) {
	_, valide := RecupererArguments([]string{"3 2 3"})

	if valide {
		t.Errorf("Les doublons doivent être refusés")
	}
}

func TestArgumentsNonNumeriques(t *testing.T) {
	_, valide := RecupererArguments([]string{"3 abc 1"})

	if valide {
		t.Errorf("Les caractères non numériques doivent être refusés")
	}
}

func TestArgumentVide(t *testing.T) {
	_, valide := RecupererArguments([]string{""})

	if valide {
		t.Errorf("Un argument vide doit être refusé")
	}
}

func TestSansArguments(t *testing.T) {
	pile, valide := RecupererArguments([]string{})

	if !valide || len(pile) != 0 {
		t.Errorf("L'absence d'arguments doit être acceptée")
	}
}
