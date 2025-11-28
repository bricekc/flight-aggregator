# Project by :
- Fatat TARRAF
- Brice KUCA
- Mohamed DIAG
- Corentin CLERO

# Flight Aggregator

Une application Go qui agrège les données de vols provenant de deux sources externes et offre une API pour consulter et trier les vols par différents critères.

## Vue d'ensemble

Flight Aggregator récupère les données de vols à partir de deux serveurs d'API externes (j-server1 et j-server2), les unifie dans une structure commune, et fournit une API REST pour accéder aux vols avec différentes options de tri.

**Critères de tri disponibles:**
- Prix (croissant/décroissant)
- Date de départ (croissant/décroissant)
- Durée du voyage (croissant/décroissant)

## Architecture

L'application suit une architecture en couches :

```
Controllers (Points d'entrée HTTP)
    ↓
Services (Logique métier)
    ↓
Repositories (Accès aux données externes)
```

- **Controllers** : Gèrent les requêtes HTTP et les réponses
- **Services** : Contiennent la logique métier (agrégation et tri des vols)
- **Repositories** : Récupèrent les données des API externes et les transforment en structures Go

## Prérequis

- **Docker** et **Docker Compose** (pour l'exécution complète)
- **Go** 1.20+ (pour le développement local)
- **GNU Make** (pour exécuter les commandes)
- `gotestsum` (pour les tests avec formattage personnalisé) : `go install gotest.tools/gotestsum@latest`

### Dépendances Go

Les dépendances principales sont listées dans `go.mod` :
- Spf13 Viper (gestion des variables d'environnement)
- Stretchr Testify (framework de test)

## Installation et configuration

### 1. Cloner le projet

```bash
git clone git@github.com:bricekc/flight-aggregator.git
git clone https://github.com/bricekc/flight-aggregator.git
cd flight-aggregator
```

### 2. Configurer les variables d'environnement

Créez un fichier `.env` à la racine du projet ou utilisez les variables d'environnement du système :

```env
SERVER_PORT=3001
JSERVER1_PORT=4001
JSERVER1_NAME=j-server1
JSERVER2_PORT=4002
JSERVER2_NAME=j-server2
JSERVER1_URL=http://j-server1:4001
JSERVER2_URL=http://j-server2:4002
```

### 3. Démarrer l'application avec Docker Compose

```bash
# Construire et démarrer les conteneurs
docker compose up --build

# Ou sur Linux
docker compose -f Docker-compose.yml up --build
```

L'application démarre avec :
- **Serveur principal** : http://localhost:3001
- **j-server1** : http://localhost:4001
- **j-server2** : http://localhost:4002

## Utilisation

### Vérifier l'état du serveur

```bash
curl http://localhost:3001/api/health
```

Réponse : `200 OK`

### Récupérer les vols

```bash
curl "http://localhost:3001/api/flight?sortby=price&orderby=asc"
```

## API Endpoints

### GET `/api/health`

Vérifie la disponibilité et l'état du serveur.

**Réponse :**
```
Status: 200 OK
Body: "OK"
```

### GET `/api/flight`

Récupère la liste de tous les vols agrégés des deux sources, avec option de tri.

**Paramètres de requête :**
| Paramètre | Type | Description | Valeurs |
|-----------|------|-------------|---------|
| `sortby` | string | Critère de tri | `price`, `departureDate`, `travelTime` |
| `orderby` | string | Ordre de tri | `asc` (croissant), `desc` (décroissant) |

**Exemple de requête :**
```bash
# Trier par prix en ordre croissant
curl "http://localhost:3001/api/flight?sortby=price&orderby=asc"

# Trier par date de départ en ordre décroissant
curl "http://localhost:3001/api/flight?sortby=departureDate&orderby=desc"

# Trier par durée du voyage
curl "http://localhost:3001/api/flight?sortby=travelTime&orderby=asc"
```

## Tests

### Exécuter tous les tests

```bash
make test
```

Cette commande exécute les tests avec `gotestsum` pour une meilleure lisibilité du résultat.

### Exécuter les tests directement

```bash
go test ./...
```

### Tests couverts

- **Tri par prix** : Vérification que les vols sont correctement ordonnés par tarif
- **Tri par date de départ** : Vérification que les vols sont correctement ordonnés par heure de départ
- **Tri par durée** : Vérification que les vols sont correctement ordonnés par durée du voyage
- **Service de vol** : Tests unitaires avec mocks des repositories
