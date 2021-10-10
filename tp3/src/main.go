package main

import (
	"fmt"
)

//fonction de service
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

//Vérifier si un cycle peut être construit après avoir ajouté un nouvel agent dans la liste
func identifyCycle(listeAgent []Agent, newAg Agent) (bool, []Agent) {
	nbA := len(listeAgent)
	existeCycle := false
	//listeAgent est de la forme [a1,b1,a2,b2,a3,b3,a4......], on va donc vérifier si newAg=listeAgent[i] avec i un nombre pair
	var i int
	for i = 0; i < nbA; i = i + 2 {
		if Equal(newAg, listeAgent[i]) {
			existeCycle = true
			break
		}
	}
	if existeCycle {
		return true, listeAgent[i:] //retourner le sous-slice qui contient le cycle complet
	} else {
		return false, nil
	}
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
					fmt.Println(As[i], " échange avec ", As[j])
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
	couple = make(map[AgentID]AgentID) //couple formé à retourner

	//initialiser un Map qui contient tous les disposants d'un proposant qui ne lui ont pas déjà refu
	disposantRestant := make(map[AgentID][]AgentID)
	//initialiser deux Maps [Agent]bool indiquant si l'agent a ou b est apparié ou pas
	apparieA := make(map[AgentID]bool)
	for _, a := range ag1 {
		apparieA[a.ID] = false
		disposantRestant[a.ID] = a.Prefs
	}
	apparieB := make(map[AgentID]bool)
	for _, b := range ag2 {
		apparieB[b.ID] = false
	}
	nonAppa, hommeCelibataire := nonApparie(apparieA)

	//tant que'il y a d'homme célibataire non apparié
	for nonAppa {
		//b = femme préférée de l'hommeCelibataire parmi celles à qui il ne s'est pas déjà proposé
		b := disposantRestant[hommeCelibataire][0]
		disposantRestant[hommeCelibataire] = disposantRestant[hommeCelibataire][1:] //retirer b dans la liste de disposant restant de l'hommeCelibataire
		//si b est célibataire, hommeCelibataire et b forment un couple
		if !apparieB[b] {
			couple[b] = hommeCelibataire
			apparieB[b] = true
			apparieA[hommeCelibataire] = true
		} else { //sinon un couple(a', b) existe
			bPrefHommeCeli, _ := findAgentByID(ag2, b).Prefers(findAgentByID(ag1, hommeCelibataire), findAgentByID(ag1, couple[b])) //si b préfère hommecelibataire à a'
			if bPrefHommeCeli {
				//b et hommeCelibataire forment un nouveau couple, hommeCelibataire devient apparié, a' devient non apparié
				apparieA[hommeCelibataire] = true
				apparieA[couple[b]] = false
				couple[b] = hommeCelibataire
			}
		}
		nonAppa, hommeCelibataire = nonApparie(apparieA)
	}
	return couple
}

func TopTradingCycles(ag1 []Agent, ag2 []Agent) (couple map[AgentID]AgentID) {
	nbA := len(ag1)
	nbB := len(ag2)
	if nbA != nbB {
		panic("A et B ne sont pas de même taille !")
	}
	couple = make(map[AgentID]AgentID) //couple formé à retourner
	//initialiser deux Maps [Agent]bool indiquant si l'agent a ou b est apparié ou pas
	apparieA := make(map[AgentID]bool)
	for _, a := range ag1 {
		apparieA[a.ID] = false
	}
	apparieB := make(map[AgentID]bool)
	for _, b := range ag2 {
		apparieB[b.ID] = false
	}
	nonAppa, hommeCelibataire := nonApparie(apparieA)
	//tant qu'il reste des agents non appariés
	for nonAppa {
		listeAgent := make([]Agent, 0)
		var cycle []Agent
		existeCycle := false
		//tant que l'on n'a pas encore construit de cycle
		for !existeCycle {
			var bPrefere Agent
			for _, bID := range findAgentByID(ag1, hommeCelibataire).Prefs {
				if !apparieB[bID] {
					bPrefere = findAgentByID(ag2, bID)
					break
				}
			}
			//listeAgent est de la forme [a1,b1,a2,b2....]
			listeAgent = append(listeAgent, findAgentByID(ag1, hommeCelibataire))
			listeAgent = append(listeAgent, bPrefere)
			var aPrefereDeB Agent
			//agent du groupe B qui a reçu une offre pointe vers son agent libre préféré du groupe A
			for _, aID := range bPrefere.Prefs {
				if !apparieA[aID] {
					aPrefereDeB = findAgentByID(ag1, aID)
					break
				}
			}
			existeCycle, cycle = identifyCycle(listeAgent, aPrefereDeB)
			//si le cycle est bien construit => apparier chaque agent du cycle avec l'agent qu'il pointe
			if existeCycle {
				nbC := len(cycle)
				var i int
				for i = 0; i < nbC; i = i + 2 {
					apparieA[cycle[i].ID] = true
					apparieB[cycle[i+1].ID] = true
					couple[cycle[i+1].ID] = cycle[i].ID
				}
			} else { //sinon l'ajouter dans listeAgent puis continuer à construire le cycle
				hommeCelibataire = aPrefereDeB.ID
			}
		}
		nonAppa, hommeCelibataire = nonApparie(apparieA)
	}
	return couple
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
		"Victor",
		"Antoine",
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
		"Lisa",
		"Anne",
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
	fmt.Println("*** AD - Acceptation Différée (a.k.a. Gale & Shapley, 1962) ***")
	coupleStable := GaleShapley(poolA, poolB)
	fmt.Println(coupleStable)
	fmt.Println(dynamiqueLibre(poolA, poolB, coupleStable))
	fmt.Println("*** Top Trading Cycles ***")
	coupleInstableParTTC := TopTradingCycles(poolA, poolB)
	fmt.Println(coupleInstableParTTC)
	fmt.Println(dynamiqueLibre(poolA, poolB, coupleInstableParTTC))
}
