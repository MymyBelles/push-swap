# Push-Swap

## Présentation

Push-Swap est un projet d'algorithmie développé en Go.

L'objectif est de trier une liste d'entiers à l'aide de deux piles
et d'un ensemble d'opérations autorisées, tout en minimisant
le nombre de mouvements.

## Objectifs pédagogiques

- Comprendre le fonctionnement des piles.
- Manipuler les slices en Go.
- Développer des algorithmes de tri.
- Optimiser le nombre de mouvements.
- Gérer les erreurs et les arguments.
- Développer des tests unitaires.

## Les deux programmes

### Push-Swap

Génère les instructions nécessaires pour trier la pile A
dans l'ordre croissant.

### Checker

Lit et exécute les instructions reçues afin de vérifier
si la pile A est triée et si la pile B est vide.

Il retourne :

- OK : le tri est correct.
- KO : le tri est incorrect.
- Error : les arguments ou les instructions sont invalides.

## Les 11 opérations


| sa : Échange les deux premiers éléments de A 
| sb : Échange les deux premiers éléments de B 
| ss : Exécute sa et sb simultanément 
| pa : Déplace le sommet de B vers A 
| pb : Déplace le sommet de A vers B 
| ra : Effectue une rotation de A 
| rb : Effectue une rotation de B 
| rr : Effectue une rotation des deux piles 
| rra : Effectue une rotation inverse de A 
| rrb : Effectue une rotation inverse de B 
| rrr : Effectue une rotation inverse des deux piles 

## Stratégie algorithmique

Le programme sélectionne automatiquement une méthode de tri
en fonction du nombre d'éléments.

- 2 éléments : triDeux.
- 3 éléments : triTrois.
- 4 éléments : triQuatre.
- 5 éléments : triCinq.
- Plus de 5 éléments : triBlocs.

Le tri par blocs repose sur les étapes suivantes :

1. Création d'une copie triée des valeurs.
2. Attribution d'un rang à chaque valeur.
3. Transfert progressif des éléments vers la pile B.
4. Recherche du maximum dans la pile B.
5. Retour des éléments vers la pile A en optimisant les rotations.

La taille des blocs est adaptée au nombre d'éléments.

## Compilation

```bash
go build -o push-swap .
go build -o checker ./cmd/checker
```

## Utilisation

### Push-Swap

```bash
./push-swap "2 1 3"
```

Résultat :

```text
sa
```

### Checker

```bash
printf "sa\n" | ./checker "2 1 3"
```

Résultat :

```text
OK
```

### Utilisation combinée

```bash
ARG="3 5 2 1 4"
./push-swap "$ARG" | ./checker "$ARG"
```

Résultat :

```text
OK
```

### Compter les mouvements

```bash
./push-swap "$ARG" | wc -l
```

## Gestion des erreurs

Le programme vérifie :

- Les arguments non numériques.
- Les doublons.
- Les arguments vides.
- Les dépassements de capacité du type int.
- Les instructions inexistantes ou mal formatées dans Checker.

Les erreurs sont affichées sur la sortie d'erreur standard.

## Organisation du projet

```text
push-swap/
├── cmd/
│   └── checker/
│       ├── main.go
│       ├── push.go
│       ├── reverse.go
│       ├── rotate.go
│       ├── sorted.go
│       ├── swap.go
│       ├── validation.go
│       └── fichiers de tests
├── main.go
├── chunks.go
├── push.go
├── reverse.go
├── rotate.go
├── sort.go
├── sorted.go
├── swap.go
├── validation.go
├── fichiers de tests
├── go.mod
├── .gitignore
└── README.md
```

## Performances observées

| Nombre d'éléments | Nombre de mouvements | Checker |
|-------------------|----------------------|---------|
| 100 | 664 | OK |
| 500 | 5711 | OK |

Ces résultats correspondent à des essais sur des listes aléatoires.
Le nombre de mouvements peut varier selon l'ordre initial des valeurs.

## Tests unitaires

Le projet comporte 49 tests unitaires :

- 27 tests pour Push-Swap.
- 22 tests pour Checker.

Les tests couvrent les opérations, la validation des arguments,
la vérification du tri et les algorithmes de tri.

Pour exécuter les tests :

```bash
go test -v ./...
```

Pour analyser le code :

```bash
go vet ./...
```

Pour formater le code :

```bash
go fmt ./...
```

## Technologies utilisées

- Go (Golang)
- Visual Studio Code
- Git
- GitHub
- claude code

Projet d'algorithmie : Push-Swap.