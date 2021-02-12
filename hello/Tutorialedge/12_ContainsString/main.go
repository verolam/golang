package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	fmt.Print("Entrez votre email : ")
	scanner.Scan()                      // lancement du scanner
	entreeUtilisateur := scanner.Text() // stockage du résultat du scanner dans une variable
	fmt.Println(entreeUtilisateur)

	// var chaine = "vlamrani@yahoo.com"
	containsChar := strings.Contains(entreeUtilisateur, "@")
	if containsChar {
		fmt.Println("Email correct")
	} else {
		fmt.Println("Email incorrect")
	}

}
