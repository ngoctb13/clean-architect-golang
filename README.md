## HOW TO START

Start infra - local environment

```bash
./scripts/bin.sh infra up
```

Start API server (include migration) - local environment

```bash
./scripts/bin.sh api start
```

## CLI (Optional)

Migrate(up) business model

```bash
./scripts/bin.sh api migrate up
```

Migrate(down) business model

```bash
./scripts/bin.sh api migrate down
```

## API

Auth API:

```bash
localhost:8090/v1/register
```

```bash
localhost:8090/v1/login
```

New API:

```bash
localhost:8090/v1/news
```