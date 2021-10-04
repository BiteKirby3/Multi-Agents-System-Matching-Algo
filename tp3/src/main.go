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

func findAgentByID(ags []Agent, id AgentID) Agent {
	var ag Agent
	for _, a := range ags {
		if a.ID == id {
			ag = a
		}
	}
	return ag
}

func paireCritique(a1 Agent, b1 Agent, a2 Agent, b2 Agent) bool {
	a1Prefb2, _ := a1.Prefers(b2, b1)
	b2Prefa1, _ := b2.Prefers(a1, a2)
	return a1Prefb2 && b2Prefa1
}

func nonApparie(appaA map[AgentID]bool) (bool, AgentID) {
	for key, val := range appaA {
		if !val {
			return true, key
		}
	}
	return false, ""
}

func Boston(ag1 []Agent, ag2 []Agent) (couple map[AgentID]AgentID) {
	nbA := len(ag1)
	nbB := len(ag2)
	if nbA != nbB {
		panic("A et B ne sont pas de même taille !")
	}
	couple = make(map[AgentID]AgentID)
	//initialiser deux Maps [Agent]bool indiquant si l'agent a ou b est apparié ou pas
	apparieA := make(map[AgentID]bool)
	for _, a := range ag1 {
		apparieA[a.ID] = false
	}
	apparieB := make(map[AgentID]bool)
	for _, b := range ag2 {
		apparieB[b.ID] = false
	}
	//dans la ième itération
	for i := 0; i < nbA; i++ {
		//chaque proposant propose s'il n'est pas encore apparié
		for _, a := range ag1 {
			if !apparieA[a.ID] {
				//chaque disposant réacte
				for _, b := range ag2 {
					//a propose à son ième choix b
					if a.Prefs[i] == b.ID {
						//chaque disposant accepte son proposant préféré
						if !apparieB[b.ID] {
							_, ok := couple[b.ID]
							if !ok {
								couple[b.ID] = a.ID
							} else {
								pref, _ := findAgentByID(ag2, b.ID).Prefers(a, findAgentByID(ag1, couple[b.ID]))
								if pref {
									couple[b.ID] = a.ID
								}
							}
						}
					}
				}
			}
		}
		//les proposants et disposants appariés sont retirés à la fin d'itération courante(mettre la valeur à true dans le map apparieA et apparieB)
		//si les couples sont tous formés, on renvoie le résultat directement
		if len(couple) == nbA {
			return couple
		}
		for bID, aID := range couple {
			apparieA[aID] = true
			apparieB[bID] = true
		}
	}
	return couple
}

//l’algorithme de dynamique libre consistant à faire se rencontrer les couples et à résoudre itérativement les instabilités.
func dynamiqueLibre(ag1 []Agent, ag2 []Agent, couple map[AgentID]AgentID) map[AgentID]AgentID {
	//transformer map en 2 arrays
	As := make([]AgentID, 0)
	Bs := make([]AgentID, 0)
	for bID, aID := range couple {
		As = append(As, aID)
		Bs = append(Bs, bID)
	}
	//boucle do while : tant qu' il existe une paire critique
	for {
		existePaireCritique := false
		for j := 0; j < len(couple); j++ {
			for i := 0; i < len(couple); i++ {
				//tant qu'il existe paire critique, on échange leur partenaire
				if paireCritique(findAgentByID(ag1, As[j]), findAgentByID(ag2, Bs[j]), findAgentByID(ag1, As[i]), findAgentByID(ag2, Bs[i])) {
					temp := As[i]
					As[i] = As[j]
					As[j] = temp
					existePaireCritique = true
					break //recommencer les comparaisons
				}
			}
			if existePaireCritique {
				break
			}
		}
		//tant qu'il n'y a pas de paire critique, on sort de la boucle
		if !existePaireCritique {
			fmt.Println("log : couple stable ")
			//mettre à jour couple à partir de As et Bs
			for i := 0; i < len(couple); i++ {
				couple[Bs[i]] = As[i]
			}
			break
		}
	}
	return couple
}

func GaleShapley(ag1 []Agent, ag2 []Agent) (couple map[AgentID]AgentID) {
	nbA := len(ag1)
	nbB := len(ag2)
	if nbA != nbB {
		panic("A et B ne sont pas de même taille !")
	}
	couple = make(map[AgentID]AgentID)
	//initialiser deux Maps [Agent]bool indiquant si l'agent a ou b est apparié ou pas
	apparieA := make(map[AgentID]bool)
	for _, a := range ag1 {
		apparieA[a.ID] = false
	}
	apparieB := make(map[AgentID]bool)
	for _, b := range ag2 {
		apparieB[b.ID] = false
	}
	nonApparie, hommeCelibataire := nonApparie(apparieA)
	rejet := make(map[AgentID][]AgentID)
	for nonApparie {

		nonApparie, hommeCelibataire = nonApparie(apparieA)
	}
	return
}

func main() {
	Anames := [...]string{
		"Khaled",
		"Sylvain",
		"Emmanuel",
		"Bob",
		"Lucas",
		"George",
		"Léo",
		"Théo",
		"Léon",
		"Louis",
		"Pierre",
	}
	Bnames := [...]string{
		"Nathalie",
		"Annaïck",
		"Brigitte",
		"Camille",
		"Léa",
		"Louise",
		"Anna",
		"Marie",
		"Sophie",
		"Nina",
		"Jeanne",
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

	fmt.Println("*** AI - Acceptation Immédiate (a.k.a. Boston) ***")
	coupleInstable := Boston(poolA, poolB)
	fmt.Println(coupleInstable)
	fmt.Println("*** DL - Dynamique Libre ***")
	fmt.Println(dynamiqueLibre(poolA, poolB, coupleInstable))

}
