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
- Organiser et maintenir un projet informatique.
- Utiliser Git et GitHub pour le versionnement.

## Les deux programmes

### Push-Swap

Génère les instructions nécessaires pour trier la pile A
dans l'ordre croissant.

La pile B est initialement vide et sert de pile auxiliaire.

À la fin du tri :

- La pile A doit être triée dans l'ordre croissant.
- La pile B doit être vide.
- Le nombre de mouvements doit être minimisé/ optimisé.

### Checker

Lit et exécute les instructions reçues afin de vérifier
si la pile A est triée et si la pile B est vide.

| Résultat | Signification |
|----------|---------------|
| OK | Le tri est correct et la pile B est vide. |
| KO | La pile A n'est pas triée ou la pile B n'est pas vide. |
| Error | Les arguments ou les instructions sont invalides. |

## Les 11 opérations

| Opération | Description |
|-----------|-------------|
| `sa` | Échange les deux premiers éléments de A. |
| `sb` | Échange les deux premiers éléments de B. |
| `ss` | Exécute sa et sb simultanément. |
| `pa` | Déplace le sommet de B vers A. |
| `pb` | Déplace le sommet de A vers B. |
| `ra` | Déplace le premier élément de A en dernière position. |
| `rb` | Déplace le premier élément de B en dernière position. |
| `rr` | Exécute ra et rb simultanément. |
| `rra` | Déplace le dernier élément de A en première position. |
| `rrb` | Déplace le dernier élément de B en première position. |
| `rrr` | Exécute rra et rrb simultanément. |

## Stratégie algorithmique

Le programme sélectionne automatiquement une méthode de tri
en fonction du nombre d'éléments.

| Nombre d'éléments | Fonction utilisée |
|-------------------|-------------------|
| 0 ou 1 | Aucun mouvement nécessaire |
| 2 | `triDeux()` |
| 3 | `triTrois()` |
| 4 | `triQuatre()` |
| 5 | `triCinq()` |
| Plus de 5 | `triBlocs()` |

### Tri par blocs

Le tri par blocs repose sur les étapes suivantes :

1. Création d'une copie triée des valeurs.
2. Attribution d'un rang à chaque valeur.
3. Transfert progressif des éléments vers la pile B.
4. Recherche du maximum dans la pile B.
5. Retour des éléments vers la pile A en optimisant les rotations.

La taille des blocs est adaptée au nombre d'éléments.

Cette stratégie vise à réduire le nombre de mouvements
nécessaires pour trier les grandes listes.

## Prérequis

- Go installé sur l'ordinateur.
- Git pour récupérer et versionner le projet.
- Un terminal pour compiler et exécuter les programmes.

Le projet utilise uniquement les packages de la bibliothèque
standard de Go.

## Compilation

Depuis la racine du projet :

```bash
go build -o push-swap .
go build -o checker ./cmd/checker
```

Ces commandes génèrent deux exécutables :

- `push-swap`
- `checker`

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
ARG="3 5 2 1 4"
./push-swap "$ARG" | wc -l
```

## Gestion des erreurs

Le programme vérifie :

- Les arguments non numériques.
- Les doublons.
- Les arguments vides.
- Les dépassements de capacité du type int.
- Les instructions inexistantes dans Checker.
- Les instructions mal formatées dans Checker.

Les erreurs sont affichées sur la sortie d'erreur standard.

En l'absence d'arguments, les programmes ne produisent
aucune sortie.

## Organisation du projet

```text
push-swap/
├── cmd/
│   └── checker/
│       ├── main.go
│       ├── operations.go
│       ├── validation.go
│       └── main_test.go
├── main.go
├── operations.go
├── sort.go
├── validation.go
├── main_test.go
├── go.mod
├── .gitignore
└── README.md
```

Le projet est organisé en 12 fichiers, hors exécutables
et fichiers internes de Git.

### Fichiers de Push-Swap

| Fichier | Rôle |
|---------|------|
| `main.go` | Point d'entrée et sélection de l'algorithme. |
| `operations.go` | Implémentation des 11 opérations autorisées. |
| `sort.go` | Algorithmes de tri et vérification du tri. |
| `validation.go` | Récupération et validation des arguments. |
| `main_test.go` | Regroupement des 27 tests unitaires. |

### Fichiers du Checker

| Fichier | Rôle |
|---------|------|
| `cmd/checker/main.go` | Lecture et exécution des instructions. |
| `cmd/checker/operations.go` | Exécution des 11 opérations sans affichage des mouvements. |
| `cmd/checker/validation.go` | Validation des arguments et vérification du tri. |
| `cmd/checker/main_test.go` | Regroupement des 22 tests unitaires. |

### Fichiers complémentaires

| Fichier | Rôle |
|---------|------|
| `go.mod` | Configuration du module Go. |
| `.gitignore` | Exclusion des exécutables et fichiers temporaires de Git. |
| `README.md` | Documentation du projet. |

## Performances observées

| Nombre d'éléments | Nombre de mouvements | Checker |
|-------------------|----------------------|---------|
| 100 | 664 | OK |
| 500 | 5711 | OK |

Ces résultats correspondent à des essais sur des listes aléatoires.

Le nombre de mouvements peut varier selon l'ordre initial
des valeurs.

## Tests unitaires

Le projet comporte 49 tests unitaires.

| Programme | Nombre de tests |
|-----------|-----------------|
| Push-Swap | 27 |
| Checker | 22 |
| Total | 49 |

Les tests couvrent :

- Les opérations Swap.
- Les opérations Push.
- Les opérations Rotate.
- Les opérations Reverse Rotate.
- La validation des arguments.
- La vérification du tri.
- Les algorithmes de tri.

### Exécuter les tests

```bash
go test -v ./...
```

### Exécuter les tests sans utiliser le cache

```bash
go test -count=1 ./...
```

### Analyser le code

```bash
go vet ./...
```

### Formater le code

```bash
go fmt ./...
```

## Technologies et outils utilisés

- Go (Golang)
- Visual Studio Code
- Git
- GitHub
- Claude Code

---

Projet d'algorithmie — Push-Swap