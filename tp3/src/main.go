package main

import (
	"fmt"
)

func PrintAg(ag Agent) {
	fmt.Println(ag)
}

func PrintInt(i int) {
	fmt.Println(i)
}

func Boston(ag1 []Agent, ag2 []Agent) (couple map[Agent]Agent) {
	couple = make(map[Agent]Agent)
	nb := len(ag1)
	for nb > 0 {
		for _, b := range ag2 {
			for _, a := range ag1 {
				_, ok := couple[b]
				if a.Prefs[0] == b.ID && !ok {
					couple[b] = a
				} else if a.Prefs[0] == b.ID && ok {
					prefer, _ := b.Prefers(a, couple[b])
					if prefer {
						couple[b] = a
					}
				}
			}
		}
		//gérer les nouvo ag1 et ag2 (retirer les couples)
		//update
		for key, val := range couple {
			for indice, a := range ag1 {
				if Equal(val, a) {
					removeAgent(indice, ag1)
				}
				for ipref, apref := range a.Prefs {
					if apref == key.ID {
						removeAgentID(ipref, a.Prefs)
					}
				}
			}
			for indiceb, b := range ag2 {
				if Equal(key, b) {
					removeAgent(indiceb, ag1)
				}
				for ipref, bpref := range b.Prefs {
					if bpref == val.ID {
						removeAgentID(ipref, b.Prefs)
					}
				}
			}

		}
		nb = len(ag1)
	}

	return
}

func removeAgent(i int, ag []Agent) []Agent {
	return append(ag[:i], ag[i+1:]...)
}

func removeAgentID(i int, ag []AgentID) []AgentID {
	return append(ag[:i], ag[i+1:]...)
}

/*
func removeAgent(i int, ag []Agent) {
	copy(ag[i:], copy[i+1:])
	ag[len(ag)-1] = nil
	ag = ag[len(ag)-1]

}

func removeAgentID(i int, ag []AgentID) {
	copy(ag[i:], copy[i+1:])
	ag[len(ag)-1] = nil
	ag = ag[len(ag)-1]

}
*/

func main() {
	Anames := [...]string{
		"Khaled",
		"Sylvain",
		"Emmanuel",
		"Bob",
	}
	Bnames := [...]string{
		"Nathalie",
		"Annaïck",
		"Brigitte",
		"Camille",
	}

	poolA := make([]Agent, 0, len(Anames))
	poolB := make([]Agent, 0, len(Bnames))

	groupA_prefix := "a"
	groupB_prefix := "b"

	prefsA := make([]AgentID, len(Anames))
	prefsB := make([]AgentID, len(Bnames))
	for i := 0; i < len(Anames); i++ {
		prefsA[i] = AgentID(groupA_prefix + fmt.Sprintf("%d", i))
	}

	for i := 0; i < len(Bnames); i++ {
		prefsB[i] = AgentID(groupB_prefix + fmt.Sprintf("%d", i))
	}

	for i := 0; i < len(Anames); i++ {
		prefs := RandomPrefs(prefsB)
		a := Agent{prefsA[i], Anames[i], prefs}
		poolA = append(poolA, a)
	}

	for i := 0; i < len(Bnames); i++ {
		prefs := RandomPrefs(prefsA)
		b := Agent{prefsB[i], Bnames[i], prefs}
		poolB = append(poolB, b)
	}

	for _, a := range poolA {
		fmt.Println(a)
	}

	for _, b := range poolB {
		fmt.Println(b)
	}

	fmt.Println(Boston(poolA, poolB))

}
