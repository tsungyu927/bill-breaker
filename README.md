# Bill Breaker

## Overview

**Bill Breaker** help you easily solves the trouble of splitting money with your friends. Allowing you to focus more on other things.

### Scenario

- Traveling with friends
- Renting house together
- ...etc

## API

[Swagger](localhost:8080/swagger/index.html)

## Features

🚀 Anonymous (Store uuid in localStorage and use it as userId & deviceId)

🚀 Grouping (Create new book)

🚀 Create cost (card)

🚀 Drag & Drop 【cost card】

🚀 Split cost automatically

### Phase 2

🚩 Identity authentication (TBC Line Login、Google…?)

🚩 Support LineBot

🚩 Notification (remind to pay)

🚩 History stats

## User flow

```mermaid
flowchart LR
    A[Enter Bill Breaker] --> B{Check userId in localstorage}
    B -->|Exist| C[To book list page]
    B -->|Not Exist| D[To create user page]
```
