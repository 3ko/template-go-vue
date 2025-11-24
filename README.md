# Mon Projet – Vue 3 + Go (Gin) + PostgreSQL + Zitadel

Projet généré automatiquement : backend Go + frontend Vue + Docker + Helm.
Voir les fichiers pour les détails.

## Démarrage rapide en développement

### Avec Docker

```bash
docker compose -f docker-compose.dev.yml up --build
```

Cette commande démarre PostgreSQL et l'API Go avec hot-reload (Air) tout en servant le bundle Vue précompilé.

### Sans Docker

Backend (Go) :

```bash
go run ./cmd/api
```

Assurez-vous d'avoir défini les variables d'environnement attendues (`PORT`, `DB_*`, `ZITADEL_*`).

Frontend (Vue 3 / Vite) :

```bash
cd client
npm install   # première fois uniquement
npm run dev
```

Le serveur Vite est accessible sur le port indiqué par la sortie (`http://localhost:5173` par défaut).

## API d'exemple (CRUD)

Une API sécurisée illustre un CRUD simple sur des utilisateurs :

- `GET /api/secure/users` : liste les utilisateurs.
- `POST /api/secure/users` : crée un utilisateur (`email`, `name`).
- `GET /api/secure/users/:id` : retourne un utilisateur par identifiant.
- `PUT /api/secure/users/:id` : met à jour un utilisateur existant.
- `DELETE /api/secure/users/:id` : supprime un utilisateur.

Toutes les routes nécessitent l'authentification via Zitadel (middlewares configurés dans `internal/api/router.go`).

## Build & livraison du frontend

Le binaire Go sert désormais directement le bundle Vue généré par Vite :

- le `Dockerfile.api` contient une étape Node qui exécute `npm run build` dans `client/` et copie le dossier `dist/` dans l'image finale ;
- le serveur Gin expose les fichiers statiques et renvoie `index.html` pour les routes hors `/api`, ce qui évite l'usage de Nginx.

En développement comme en production, un seul conteneur `api` est nécessaire : les services `client` ont été supprimés des fichiers `docker-compose.*`. Le port 8080 sert à la fois l'API (`/api/...`) et l'interface Vue (racine et routes SPA).
