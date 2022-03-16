package adder_test

import (
    "gitlab.eemi.tech/yannis.bengueci/adder"
    "testing"
    "log"
    "os"
)

func TestAdd(t *testing.T) { // https://play.golang.org/p/ECID9_9ddJm
    checks := []struct {numbers []int; expected int; name string}{
        {[]int{}, 0, "vide"},
        {[]int{5}, 5, "une seule entrée"},
        {[]int{1, 2, 3, 4, 5}, 15, "cinq entrées"},
    }
    for _, check := range checks {
        t.Run(check.name, func(t *testing.T) {
            check := check
            t.Parallel()
            if actual := adder.Add(check.numbers...); actual != check.expected {
                t.Errorf("expected %d, got %d", check.expected, actual)
            }
        })
    }
}

func TestFactoriel(t *testing.T) {
    checks := []struct {num int; expected int; name string}{
        {-8, 0, "Error negative num"},
        {0, 1, "factoriel de 0 = 1"},
        {1, 1, "factoriel de 1 = 1"},
        {7, 5040, "factoriel de 7 = 5040"},
    }

    for _, check := range checks {
        t.Run(check.name, func(t *testing.T) {
            check := check
            t.Parallel()
            if actual := adder.Factoriel(check.num); actual != check.expected {
                t.Errorf("expected %d, got %d", check.expected, actual)
            }
        })
    }
}

func TestMain(m *testing.M) {
    log.Print("setup de paquet")
    sts := m.Run()
    log.Print("teardown/cleanup de paquet")
    os.Exit(sts)
}
