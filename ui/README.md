# Curio Dashboard UI

## Development

### Prerequisites

- [Node](https://nodejs.org/) >= v22
- [pnpm](https://pnpm.io/) >= 10

### Setup

```bash
git clone https://github.com/web3tea/curio-dashboard.git

cd curio-dashboard/ui
```

### Install dependencies

```bash
pnpm i
```

### Start the development server

```bash
pnpm dev

# or
VITE_SERVER_URL=http://your.server pnpm dev

# The default server url is http://localhost:9091/graphql
```

### Build for production

```bash
pnpm build

# or
VITE_SERVER_URL=http://your.server pnpm build

# The default server url is /graphql
```

## Structure

```
ui/
├── public/                 # Static assets
├── src/                    # Source code
│   ├── assets/             # Project assets (images, fonts, etc.)
│   ├── components/         # Vue components
│   ├── composables/        # Reusable logic (composables)
│   ├── gql/                # GraphQL queries and mutations
│   ├── plugins/            # Vue plugins
│   ├── router/             # Vue Router configuration
│   ├── stores/             # Pinia stores
│   ├── styles/             # Global styles
│   ├── views/              # Vue views (pages)
│   ├── App.vue             # Main App component
│   ├── main.ts             # Entry point
├── .eslintrc.js            # ESLint configuration
├── .prettierrc             # Prettier configuration
├── index.html              # HTML template
├── package.json            # Project dependencies and scripts
├── tsconfig.json           # TypeScript configuration
└── vite.config.ts          # Vite configuration
```
