# Documentation Bot Design Document (Go-First)

## Executive Summary

This design document outlines an AI-powered documentation bot built **initially for Go applications** with an **extensible architecture** for future multi-language support. The bot analyzes merged PRs, generates comprehensive markdown documentation, and creates new PRs for the docs team to review.

**Current Implementation:** Go APIs with Gin framework
**Future Support:** React, Vue, Python, Node.js, TypeScript, Infrastructure-as-Code, and other stacks (extensible design)

**Key Principle:** Start with Go demo, prove value, then extend to other languages.

## 1. System Overview

### 1.1 Goals (Go MVP)

- **Automated Markdown Documentation Generation**: Generate comprehensive markdown documentation from Go codebases
- **Documentation Improvement**: Analyze existing documentation and create improved versions
- **Automated PR Creation**: Submit new PRs with documentation file changes (not inline comments)
- **Multi-Format Support**: Generate README, API docs, architecture docs, godoc
- **Extensible Design**: Architecture supports future expansion to other languages
- **Docs Team Workflow**: PRs are reviewed and approved by docs team before merging
- **Verbose Output**: Generate detailed, comprehensive documentation (not minimal)
- **Smart Detection**: Only create PR if documentation actually needs updating

### 1.2 Key Features

1. **Go Code Analysis** - Parses Go files using AST (go/parser, go/ast)
2. **AI-Powered Generation** - Uses Claude Code GitHub Action to generate verbose markdown
3. **Automated PR Creation** - Creates documentation file changes and submits PRs
4. **Smart Documentation Detection** - Only creates PR if docs are missing or outdated
5. **GitHub Actions Integration** - Triggers on merged PRs automatically
6. **Docs Team Review** - All generated docs reviewed by team before merging
7. **Extensible Architecture** - Designed to add support for other languages later

### 1.3 Current & Future Support

| Status | Project Type | Documentation Generated |
|--------|--------------|-------------------------|
| **✅ Current (MVP)** | Go APIs (Gin) | API docs, architecture, data models, godoc, README |
| **🔜 Phase 2** | React/TypeScript Frontend | Component docs, props, state management, routing |
| **🔜 Phase 3** | Python Backend | FastAPI/Django docs, data models, endpoints |
| **🔜 Phase 4** | Infrastructure (Terraform) | Resource docs, deployment guides, network topology |
| **🔜 Future** | Other languages | Extensible analyzer framework |

**Note:** This demo focuses on Go. The architecture is designed to be extended to other languages by adding language-specific analyzers.

### 1.4 Workflow Overview

```
PR Merged/Push to Main → Bot Triggered → Analyze Changes → Generate Comprehensive Markdown → Create New PR
```

**Key Difference**: This bot creates **actual file changes** and submits **new PRs**, rather than posting suggestions as comments.

**Trigger Options:**
1. **Automatic**: After PR is merged to main
2. **Label-Based**: Add `docs:generate` label to trigger
3. **Manual**: Use GitHub Actions workflow_dispatch
4. **Scheduled**: Weekly cron job to keep docs fresh

---

## 2. Architecture

### 2.1 High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                      GitHub Repository                          │
│                                                                   │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐     │
│  │   Go Code    │    │     Docs     │    │   Tests      │     │
│  │   (*.go)     │    │  (*.md, doc) │    │  (*_test.go) │     │
│  └──────┬───────┘    └──────┬───────┘    └──────┬───────┘     │
│         │                   │                   │               │
└─────────┼───────────────────┼───────────────────┼───────────────┘
          │                   │                   │
          │  Push/PR          │                   │
          └───────────────────┴───────────────────┘
                              │
          ┌───────────────────▼───────────────────┐
          │     GitHub Actions Workflow            │
          │                                        │
          │  ┌──────────────────────────────┐    │
          │  │   1. Code Analysis           │    │
          │  │      - Parse Go files        │    │
          │  │      - Extract structure     │    │
          │  └───────────┬──────────────────┘    │
          │              │                        │
          │  ┌───────────▼──────────────────┐    │
          │  │   2. Documentation Bot       │    │
          │  │      - Claude API calls      │    │
          │  │      - Generate docs         │    │
          │  │      - Improve existing      │    │
          │  └───────────┬──────────────────┘    │
          │              │                        │
          │  ┌───────────▼──────────────────┐    │
          │  │   3. Swagger Generation      │    │
          │  │      - swaggo/swag           │    │
          │  │      - OpenAPI spec          │    │
          │  └───────────┬──────────────────┘    │
          │              │                        │
          │  ┌───────────▼──────────────────┐    │
          │  │   4. Quality Validation      │    │
          │  │      - Vale (prose linting)  │    │
          │  │      - markdownlint          │    │
          │  │      - godoc validation      │    │
          │  │      - Doc Detective tests   │    │
          │  └───────────┬──────────────────┘    │
          │              │                        │
          │  ┌───────────▼──────────────────┐    │
          │  │   5. Report & PR Creation    │    │
          │  │      - Generate report       │    │
          │  │      - Create/update PR      │    │
          │  │      - Add comments          │    │
          │  └──────────────────────────────┘    │
          └────────────────────────────────────────┘
```

### 2.2 Component Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                     Documentation Bot Core                       │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  ┌────────────────────────────────────────────────────────┐    │
│  │              Code Analysis Layer                        │    │
│  ├────────────────────────────────────────────────────────┤    │
│  │  - AST Parser (go/parser, go/ast)                      │    │
│  │  - Function/Type Extractor                              │    │
│  │  - Existing Doc Analyzer                                │    │
│  │  - Complexity Analyzer                                  │    │
│  └────────────────┬───────────────────────────────────────┘    │
│                   │                                              │
│  ┌────────────────▼───────────────────────────────────────┐    │
│  │           AI Documentation Engine                       │    │
│  ├────────────────────────────────────────────────────────┤    │
│  │  - Claude API Client                                    │    │
│  │  - Prompt Engineering Templates                         │    │
│  │  - Context Builder                                      │    │
│  │  - Response Parser                                      │    │
│  │  - Doc Formatter                                        │    │
│  └────────────────┬───────────────────────────────────────┘    │
│                   │                                              │
│  ┌────────────────▼───────────────────────────────────────┐    │
│  │           Documentation Generator                       │    │
│  ├────────────────────────────────────────────────────────┤    │
│  │  - Godoc Comment Generator                              │    │
│  │  - Swagger Annotation Generator                         │    │
│  │  - Markdown Doc Generator                               │    │
│  │  - README Generator                                     │    │
│  │  - Code Example Generator                               │    │
│  └────────────────┬───────────────────────────────────────┘    │
│                   │                                              │
│  ┌────────────────▼───────────────────────────────────────┐    │
│  │          Validation & Testing Layer                     │    │
│  ├────────────────────────────────────────────────────────┤    │
│  │  - Vale Integration (prose quality)                     │    │
│  │  - Markdownlint (formatting)                            │    │
│  │  - Godoc Validator                                      │    │
│  │  - Link Checker                                         │    │
│  │  - Example Code Tester                                  │    │
│  │  - Doc Detective Integration                            │    │
│  └────────────────┬───────────────────────────────────────┘    │
│                   │                                              │
│  ┌────────────────▼───────────────────────────────────────┐    │
│  │            Output & Integration Layer                   │    │
│  ├────────────────────────────────────────────────────────┤    │
│  │  - File Writer                                          │    │
│  │  - Git Integration                                      │    │
│  │  - PR Creator                                           │    │
│  │  - Report Generator                                     │    │
│  └────────────────────────────────────────────────────────┘    │
│                                                                   │
└─────────────────────────────────────────────────────────────────┘
```

---

## 3. Workflow Diagrams

### 3.1 Documentation Bot Workflow (Mermaid)

```mermaid
flowchart TD
    Start([Developer Merges PR]) --> Trigger{GitHub Action<br/>Triggered?}
    Trigger -->|Yes| Checkout[Checkout Code<br/>fetch-depth: 0]
    Trigger -->|No| End1([End - No Action])

    Checkout --> GetPRInfo[Get PR Information<br/>- Changed files<br/>- PR title/description<br/>- Related modules]

    GetPRInfo --> ClaudeAction[Claude Code Action<br/>Analyzes codebase]

    ClaudeAction --> AnalyzeModules[Analyze Related Modules<br/>- Parse code structure<br/>- Extract API/components<br/>- Map dependencies<br/>- Read existing docs]

    AnalyzeModules --> CheckDocs{Documentation<br/>Exists?}

    CheckDocs -->|Yes| CheckCurrent{Is it<br/>up-to-date?}
    CheckDocs -->|No| GenerateDocs[Generate New Docs<br/>- API.md<br/>- ARCHITECTURE.md<br/>- COMPONENTS.md<br/>- README updates]

    CheckCurrent -->|Yes| NoChange[Create file:<br/>no_docs_needed.txt]
    CheckCurrent -->|No| UpdateDocs[Update Existing Docs<br/>- Improve clarity<br/>- Add missing sections<br/>- Update examples]

    GenerateDocs --> Validate[Basic Validation<br/>- markdownlint<br/>- Check formatting]
    UpdateDocs --> Validate

    Validate --> CreateBranch[Create Branch<br/>docs/update-pr-NUMBER]

    CreateBranch --> CommitChanges[Commit Changes<br/>with detailed message]

    CommitChanges --> CheckChanges{Files<br/>Modified?}

    CheckChanges -->|Yes| CreatePR[Create Pull Request<br/>- Title: docs update<br/>- Label: documentation<br/>- Assign: docs-team]
    CheckChanges -->|No| NoChange

    NoChange --> End2([End - No PR Created])

    CreatePR --> NotifyTeam[Notify Docs Team<br/>for review]

    NotifyTeam --> End3([End - Awaiting Review])

    style Start fill:#e1f5e1
    style End1 fill:#ffe1e1
    style End2 fill:#fff3cd
    style End3 fill:#d1ecf1
    style ClaudeAction fill:#cfe2ff
    style CreatePR fill:#d1ecf1
```

### 3.2 Decision Logic Flow

```mermaid
flowchart LR
    PR[Merged PR] --> Analyze{Analyze<br/>Changes}

    Analyze -->|Frontend| FE[Check:<br/>- Components<br/>- Routes<br/>- State mgmt]
    Analyze -->|Backend| BE[Check:<br/>- API endpoints<br/>- Models<br/>- Services]
    Analyze -->|Infra| Infra[Check:<br/>- Resources<br/>- Networks<br/>- Security]

    FE --> DocDecision{Documentation<br/>Status?}
    BE --> DocDecision
    Infra --> DocDecision

    DocDecision -->|Missing| Create[Create<br/>Documentation]
    DocDecision -->|Outdated| Update[Update<br/>Documentation]
    DocDecision -->|Current| Skip[Skip - No<br/>Changes Needed]

    Create --> NewPR[New PR<br/>Created]
    Update --> NewPR
    Skip --> NoAction[No PR<br/>Created]

    style Create fill:#d4edda
    style Update fill:#fff3cd
    style Skip fill:#f8d7da
    style NewPR fill:#d1ecf1
```

### 3.2 Testing Workflow

```
┌──────────────────────────────────────────────────┐
│           Documentation Test Suite               │
└──────────────────┬───────────────────────────────┘
                   │
      ┌────────────┼────────────┐
      │            │            │
      ▼            ▼            ▼
┌──────────┐ ┌──────────┐ ┌──────────┐
│ Syntax   │ │ Quality  │ │ Accuracy │
│ Tests    │ │ Tests    │ │ Tests    │
└────┬─────┘ └────┬─────┘ └────┬─────┘
     │            │            │
     ▼            ▼            ▼
┌──────────┐ ┌──────────┐ ┌──────────┐
│godoc     │ │Vale      │ │Doc       │
│validate  │ │lint      │ │Detective │
├──────────┤ ├──────────┤ ├──────────┤
│swagger   │ │markdown  │ │Example   │
│validate  │ │lint      │ │tests     │
├──────────┤ ├──────────┤ ├──────────┤
│link      │ │spelling  │ │API       │
│check     │ │check     │ │tests     │
└────┬─────┘ └────┬─────┘ └────┬─────┘
     │            │            │
     └────────────┼────────────┘
                  │
                  ▼
         ┌────────────────┐
         │ Aggregate      │
         │ Results        │
         └────────┬───────┘
                  │
             ┌────┴────┐
             │         │
          Pass       Fail
             │         │
             ▼         ▼
         ┌────┐   ┌────────┐
         │ ✓  │   │ Report │
         └────┘   │ Issues │
                  └────────┘
```

---

## 4. Markdown Documentation Types (Language-Agnostic)

### 4.1 Documentation Files Generated

The bot generates **verbose, comprehensive** markdown documentation tailored to your project type:

| File | Purpose | Applicable To | Verbosity Level |
|------|---------|---------------|-----------------|
| `README.md` | Project overview, quick start | All projects | High - detailed guide |
| `docs/API.md` | API endpoint documentation | Backend APIs | Very High - complete reference |
| `docs/COMPONENTS.md` | Component documentation | Frontend (React/Vue/Angular) | Very High - props, events, examples |
| `docs/ARCHITECTURE.md` | System architecture | All projects | Very High - deep dive |
| `docs/USAGE.md` | Usage examples & tutorials | All projects | High - step-by-step |
| `docs/DEPLOYMENT.md` | Deployment guide | All projects | High - environment-specific |
| `docs/DATA_MODELS.md` | Data models & schemas | Backend/Data projects | Very High - field descriptions |
| `docs/INFRASTRUCTURE.md` | Infrastructure as code docs | IaC projects (Terraform, etc.) | Very High - resource explanations |
| `docs/STATE_MANAGEMENT.md` | State management docs | Frontend (Redux, Vuex, etc.) | High - actions, reducers, store |
| `docs/TESTING.md` | Testing guide | All projects | Medium - how to run tests |

### 4.2 Project-Specific Examples

#### 4.2.1 React Frontend Project

**Files Generated:**
- `README.md` - Enhanced with component tree, routing structure
- `docs/COMPONENTS.md` - All React components with props, state, hooks
- `docs/API_INTEGRATION.md` - How frontend talks to backend
- `docs/STATE_MANAGEMENT.md` - Redux/Context/Zustand documentation
- `docs/ROUTING.md` - React Router setup and protected routes
- `docs/STYLING.md` - CSS modules, Tailwind, styled-components guide

**Example:** `docs/COMPONENTS.md` for React:

```markdown
# React Components Documentation

## Overview

This document provides comprehensive documentation for all React components in the application. Each component includes props documentation, usage examples, and implementation notes.

## Component Hierarchy

```
<App>
  ├── <Header>
  │   ├── <Logo>
  │   └── <Navigation>
  │       └── <NavItem>
  ├── <MainContent>
  │   ├── <Sidebar>
  │   └── <ContentArea>
  │       ├── <UserList>
  │       │   └── <UserCard>
  │       └── <ProductGrid>
  │           └── <ProductCard>
  └── <Footer>
```

---

## Components

### UserCard

**File:** `src/components/UserCard.tsx`

**Purpose:** Displays user information in a card format with avatar, name, email, and role.

**Props:**

| Prop | Type | Required | Default | Description |
|------|------|----------|---------|-------------|
| `user` | `User` | Yes | - | User object containing user data |
| `onEdit` | `(id: string) => void` | No | - | Callback when edit button is clicked |
| `onDelete` | `(id: string) => void` | No | - | Callback when delete button is clicked |
| `showActions` | `boolean` | No | `true` | Whether to show edit/delete actions |
| `className` | `string` | No | `''` | Additional CSS classes |

**User Type:**

```typescript
interface User {
  id: string;
  name: string;
  email: string;
  role: 'admin' | 'user';
  avatar?: string;
  createdAt: Date;
}
```

**Usage Example:**

```tsx
import { UserCard } from '@/components/UserCard';

function UserList() {
  const handleEdit = (id: string) => {
    console.log('Edit user:', id);
  };

  const handleDelete = (id: string) => {
    console.log('Delete user:', id);
  };

  return (
    <div className="grid grid-cols-3 gap-4">
      <UserCard
        user={{
          id: '123',
          name: 'John Doe',
          email: 'john@example.com',
          role: 'user',
          avatar: '/avatars/john.jpg',
          createdAt: new Date()
        }}
        onEdit={handleEdit}
        onDelete={handleDelete}
      />
    </div>
  );
}
```

**State Management:**

This component is stateless. It receives data via props and communicates user actions via callbacks.

**Styling:**

Uses Tailwind CSS utility classes with custom component-specific styles defined in `UserCard.module.css`.

**Accessibility:**

- Keyboard navigable with Tab/Enter
- ARIA labels for action buttons
- Screen reader friendly
- Focus indicators visible

**Testing:**

```tsx
// UserCard.test.tsx
import { render, screen, fireEvent } from '@testing-library/react';
import { UserCard } from './UserCard';

describe('UserCard', () => {
  const mockUser = {
    id: '123',
    name: 'John Doe',
    email: 'john@example.com',
    role: 'user' as const,
    createdAt: new Date()
  };

  it('renders user information', () => {
    render(<UserCard user={mockUser} />);

    expect(screen.getByText('John Doe')).toBeInTheDocument();
    expect(screen.getByText('john@example.com')).toBeInTheDocument();
    expect(screen.getByText('user')).toBeInTheDocument();
  });

  it('calls onEdit when edit button is clicked', () => {
    const handleEdit = jest.fn();
    render(<UserCard user={mockUser} onEdit={handleEdit} />);

    fireEvent.click(screen.getByRole('button', { name: /edit/i }));

    expect(handleEdit).toHaveBeenCalledWith('123');
  });
});
```

**Performance Considerations:**

- Memoized with `React.memo()` to prevent unnecessary re-renders
- Avatar images lazy-loaded
- Event handlers memoized with `useCallback()`

[... continues for all components ...]
```

#### 4.2.2 Python Backend (FastAPI/Django) Project

**Files Generated:**
- `README.md` - Enhanced with API overview, setup instructions
- `docs/API.md` - All REST endpoints with request/response examples
- `docs/DATA_MODELS.md` - SQLAlchemy/Django ORM models
- `docs/ARCHITECTURE.md` - Service layer, repository pattern
- `docs/AUTHENTICATION.md` - JWT/OAuth setup
- `docs/DEPLOYMENT.md` - Docker, AWS/GCP deployment

#### 4.2.3 Infrastructure (Terraform) Project

**Files Generated:**
- `README.md` - Enhanced with infrastructure overview
- `docs/INFRASTRUCTURE.md` - All Terraform resources explained
- `docs/ARCHITECTURE.md` - Network topology, security groups
- `docs/DEPLOYMENT.md` - How to apply changes safely
- `docs/VARIABLES.md` - All input variables documented
- `docs/OUTPUTS.md` - All output values explained

#### 4.2.4 Mobile (React Native) Project

**Files Generated:**
- `README.md` - Enhanced with mobile-specific setup (iOS/Android)
- `docs/SCREENS.md` - All screens and navigation
- `docs/COMPONENTS.md` - Reusable components
- `docs/API_INTEGRATION.md` - API client, auth, offline support
- `docs/STATE_MANAGEMENT.md` - Redux/MobX/Context
- `docs/NATIVE_MODULES.md` - Custom native modules docs

### 4.2 Example: API.md (Verbose Output)

Here's what the bot generates for API documentation:

```markdown
# API Documentation

Last updated: 2025-01-15
Generated by: Documentation Bot

## Overview

This document provides comprehensive documentation for all API endpoints in the application. Each endpoint includes detailed descriptions, request/response schemas, error handling, and code examples in multiple languages.

## Base URL

```
http://localhost:8080/api/v1
```

## Authentication

All API endpoints (except `/health`) require authentication via JWT tokens. Include the token in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

---

## Endpoints

### Health Check

#### `GET /api/v1/health`

Returns the health status of the API server.

**Authentication**: Not required

**Request**

No parameters required.

**Response**

Status: `200 OK`

```json
{
  "status": "healthy",
  "timestamp": "2025-01-15T10:30:00Z",
  "version": "1.0.0",
  "uptime": "24h15m30s"
}
```

**Error Responses**

| Status Code | Description | Example Response |
|-------------|-------------|------------------|
| 500 | Server error | `{"error": "internal server error"}` |

**Code Examples**

<details>
<summary>cURL</summary>

```bash
curl -X GET http://localhost:8080/api/v1/health
```
</details>

<details>
<summary>Go</summary>

```go
resp, err := http.Get("http://localhost:8080/api/v1/health")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var health HealthResponse
json.NewDecoder(resp.Body).Decode(&health)
fmt.Printf("Status: %s\n", health.Status)
```
</details>

<details>
<summary>Python</summary>

```python
import requests

response = requests.get('http://localhost:8080/api/v1/health')
health = response.json()
print(f"Status: {health['status']}")
```
</details>

---

### Create User

#### `POST /api/v1/users`

Creates a new user in the system.

**Authentication**: Required

**Request**

Headers:
```
Content-Type: application/json
Authorization: Bearer <token>
```

Body:
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "role": "user"
}
```

**Request Schema**

| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| `name` | string | Yes | 2-100 chars | User's full name |
| `email` | string | Yes | Valid email format | User's email address (must be unique) |
| `role` | string | No | Enum: user, admin | User role (default: user) |

**Response**

Status: `201 Created`

```json
{
  "id": "usr_1234567890",
  "name": "John Doe",
  "email": "john.doe@example.com",
  "role": "user",
  "created_at": "2025-01-15T10:30:00Z",
  "updated_at": "2025-01-15T10:30:00Z"
}
```

**Response Schema**

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique user identifier |
| `name` | string | User's full name |
| `email` | string | User's email address |
| `role` | string | User role (user or admin) |
| `created_at` | string (ISO 8601) | Creation timestamp |
| `updated_at` | string (ISO 8601) | Last update timestamp |

**Error Responses**

| Status Code | Description | Example Response |
|-------------|-------------|------------------|
| 400 | Invalid request body | `{"error": "invalid email format"}` |
| 401 | Unauthorized | `{"error": "missing or invalid token"}` |
| 409 | Email already exists | `{"error": "user with this email already exists"}` |
| 500 | Server error | `{"error": "internal server error"}` |

**Code Examples**

<details>
<summary>cURL</summary>

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "role": "user"
  }'
```
</details>

<details>
<summary>Go</summary>

```go
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Role  string `json:"role,omitempty"`
}

func createUser(token string) error {
    user := CreateUserRequest{
        Name:  "John Doe",
        Email: "john.doe@example.com",
        Role:  "user",
    }

    body, _ := json.Marshal(user)
    req, _ := http.NewRequest("POST", "http://localhost:8080/api/v1/users", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return fmt.Errorf("failed to create user: %s", resp.Status)
    }

    return nil
}
```
</details>

<details>
<summary>Python</summary>

```python
import requests

def create_user(token):
    url = 'http://localhost:8080/api/v1/users'
    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {token}'
    }
    data = {
        'name': 'John Doe',
        'email': 'john.doe@example.com',
        'role': 'user'
    }

    response = requests.post(url, json=data, headers=headers)

    if response.status_code == 201:
        user = response.json()
        print(f"Created user: {user['id']}")
        return user
    else:
        raise Exception(f"Failed to create user: {response.text}")
```
</details>

<details>
<summary>JavaScript</summary>

```javascript
async function createUser(token) {
    const response = await fetch('http://localhost:8080/api/v1/users', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({
            name: 'John Doe',
            email: 'john.doe@example.com',
            role: 'user'
        })
    });

    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }

    const user = await response.json();
    console.log('Created user:', user.id);
    return user;
}
```
</details>

**Business Logic Notes**

- Email addresses are automatically converted to lowercase
- Role defaults to "user" if not specified
- Admin users can create other admin users
- Email verification is sent asynchronously (does not block the response)

**Rate Limiting**

This endpoint is rate-limited to 100 requests per hour per IP address.

**Related Endpoints**

- [`GET /api/v1/users/:id`](#get-user-by-id) - Retrieve created user
- [`PUT /api/v1/users/:id`](#update-user) - Update user details
- [`DELETE /api/v1/users/:id`](#delete-user) - Delete user

---

[... similar detailed documentation for all other endpoints ...]
```

### 4.3 Example: ARCHITECTURE.md (Verbose Output)

```markdown
# System Architecture

Last updated: 2025-01-15
Generated by: Documentation Bot

## Table of Contents

1. [Overview](#overview)
2. [High-Level Architecture](#high-level-architecture)
3. [Component Details](#component-details)
4. [Data Flow](#data-flow)
5. [Technology Stack](#technology-stack)
6. [Design Decisions](#design-decisions)
7. [Security Considerations](#security-considerations)
8. [Performance Characteristics](#performance-characteristics)
9. [Deployment Architecture](#deployment-architecture)

## Overview

This document provides a comprehensive overview of the system architecture for the Doc Agent Demo application. The application follows a clean architecture pattern with clear separation of concerns between HTTP handlers, business logic, and data access layers.

### Key Architectural Principles

1. **Separation of Concerns**: Each layer has a single, well-defined responsibility
2. **Dependency Inversion**: High-level modules do not depend on low-level modules
3. **Testability**: All components are designed to be easily unit-testable
4. **Scalability**: Stateless design allows horizontal scaling
5. **Maintainability**: Clear structure makes the codebase easy to navigate and modify

## High-Level Architecture

```
┌─────────────────────────────────────────────────────────┐
│                      Client Layer                        │
│  (HTTP Clients, Web Browsers, Mobile Apps)               │
└──────────────────┬──────────────────────────────────────┘
                   │ HTTP/HTTPS
                   ▼
┌─────────────────────────────────────────────────────────┐
│                   API Gateway / Load Balancer            │
│              (nginx, AWS ALB, etc.)                      │
└──────────────────┬──────────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────────┐
│                 Application Server (Gin)                 │
│                                                          │
│  ┌──────────────────────────────────────────────────┐  │
│  │           Middleware Layer                       │  │
│  │  - Authentication                                │  │
│  │  - Logging                                       │  │
│  │  - CORS                                          │  │
│  │  - Rate Limiting                                 │  │
│  └────────────────┬─────────────────────────────────┘  │
│                   │                                      │
│  ┌────────────────▼─────────────────────────────────┐  │
│  │           HTTP Handlers (Controllers)            │  │
│  │  - Health Handler                                │  │
│  │  - User Handler                                  │  │
│  │  - Product Handler                               │  │
│  └────────────────┬─────────────────────────────────┘  │
│                   │                                      │
│  ┌────────────────▼─────────────────────────────────┐  │
│  │          Service Layer (Business Logic)          │  │
│  │  - User Service                                  │  │
│  │  - Product Service                               │  │
│  │  - Auth Service                                  │  │
│  └────────────────┬─────────────────────────────────┘  │
│                   │                                      │
│  ┌────────────────▼─────────────────────────────────┐  │
│  │         Repository Layer (Data Access)           │  │
│  │  - User Repository                               │  │
│  │  - Product Repository                            │  │
│  └────────────────┬─────────────────────────────────┘  │
└───────────────────┼──────────────────────────────────────┘
                    │
       ┌────────────┴────────────┐
       │                         │
       ▼                         ▼
┌──────────────┐        ┌──────────────┐
│   Database   │        │    Cache     │
│  (PostgreSQL)│        │   (Redis)    │
└──────────────┘        └──────────────┘
```

[... continues with detailed explanations of each component, data flows, etc. ...]
```

### 4.4 README.md Improvements

The bot updates README.md sections with verbose content:

```markdown
# Doc Agent Demo

A demonstration project for an AI-powered documentation bot that automatically generates and improves documentation for Go applications.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
  - [Running Locally](#running-locally)
  - [API Endpoints](#api-endpoints)
  - [Testing](#testing)
- [Documentation](#documentation)
- [Development](#development)
  - [Project Structure](#project-structure)
  - [Building](#building)
  - [Testing](#testing-1)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

## Overview

[Verbose, detailed overview of the project with motivation, use cases, and benefits...]

## Features

- ✅ **RESTful API** - Full CRUD operations for users and products
- ✅ **Authentication** - JWT-based authentication system
- ✅ **API Documentation** - Auto-generated Swagger/OpenAPI docs
- ✅ **Comprehensive Logging** - Structured logging with request tracing
- ✅ **Testing** - Unit and integration tests with >80% coverage
- ✅ **CI/CD** - Automated testing and deployment via GitHub Actions
- ✅ **Documentation Bot** - AI-powered documentation generation

[... detailed explanation of each feature ...]

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed on your system:

- **Go 1.23 or higher** - [Download and install Go](https://go.dev/doc/install)
  ```bash
  go version  # Should output go1.23 or higher
  ```

- **Git** - [Install Git](https://git-scm.com/downloads)
  ```bash
  git --version
  ```

- **Docker** (optional, for containerized deployment)
  ```bash
  docker --version
  ```

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/yourorg/doc-agent-demo.git
   cd doc-agent-demo
   ```

2. **Install dependencies**

   ```bash
   go mod download
   go mod verify
   ```

3. **Set up environment variables**

   Create a `.env` file in the project root:

   ```bash
   cp .env.example .env
   ```

   Edit `.env` and configure the following variables:

   ```env
   # Server Configuration
   PORT=8080
   ENV=development

   # API Keys
   ANTHROPIC_API_KEY=your_api_key_here

   # Database (if using)
   DATABASE_URL=postgresql://user:pass@localhost:5432/docagent
   ```

[... continues with detailed step-by-step instructions ...]
```

---

## 5. Technology Stack

### 4.1 Core Technologies

| Component | Technology | Purpose |
|-----------|-----------|---------|
| **Language** | Go 1.23+ | Main application language |
| **AI Engine** | Claude API (Sonnet 4.5) | Documentation generation |
| **Web Framework** | Gin | REST API framework |
| **API Documentation** | swaggo/swag | OpenAPI/Swagger generation |
| **CI/CD** | GitHub Actions | Automation pipeline |

### 4.2 Documentation Tools

| Tool | Purpose | Integration |
|------|---------|-------------|
| **Vale** | Prose linting, style checking | CI/CD pipeline |
| **markdownlint** | Markdown formatting validation | Pre-commit hook + CI |
| **Doc Detective** | Documentation accuracy testing | CI/CD pipeline |
| **godoc** | Go documentation generation | Native Go tool |
| **swagger-ui** | API documentation viewer | Web interface |

### 4.3 Development Tools

| Tool | Purpose |
|------|---------|
| **go/parser** | AST parsing |
| **go/ast** | Syntax tree analysis |
| **gh CLI** | GitHub API interaction |
| **git** | Version control |

---

## 5. Sample Go Application Structure

### 5.1 Project Layout

```
doc-agent-demo/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── handlers/                   # HTTP handlers
│   │   ├── health.go
│   │   ├── user.go
│   │   └── product.go
│   ├── models/                     # Data models
│   │   ├── user.go
│   │   └── product.go
│   ├── services/                   # Business logic
│   │   ├── user_service.go
│   │   └── product_service.go
│   └── middleware/                 # HTTP middleware
│       ├── auth.go
│       └── logger.go
├── pkg/
│   └── docbot/                     # Documentation bot package
│       ├── analyzer.go             # Code analysis
│       ├── claude.go               # Claude API client
│       ├── generator.go            # Doc generation
│       └── validator.go            # Doc validation
├── docs/                           # Generated documentation
│   ├── swagger.json
│   ├── swagger.yaml
│   └── docs.go
├── scripts/
│   └── generate-docs.sh            # Documentation generation script
├── .github/
│   └── workflows/
│       ├── docs.yml                # Documentation workflow
│       └── tests.yml               # Test workflow
├── .vale/                          # Vale configuration
│   ├── styles/
│   └── vocab/
├── tests/
│   ├── integration/
│   └── documentation/              # Documentation tests
│       └── doc-detective.json
├── .markdownlint.json              # Markdownlint config
├── .vale.ini                       # Vale config
├── go.mod
├── go.sum
└── README.md
```

### 5.2 Sample API Endpoints

```go
// Simple REST API structure
GET    /api/v1/health              # Health check
GET    /api/v1/users               # List users
GET    /api/v1/users/:id           # Get user by ID
POST   /api/v1/users               # Create user
PUT    /api/v1/users/:id           # Update user
DELETE /api/v1/users/:id           # Delete user
GET    /api/v1/products            # List products
GET    /api/v1/products/:id        # Get product by ID
POST   /api/v1/products            # Create product
GET    /swagger/*any               # Swagger UI
```

---

## 6. Implementation Details

### 6.1 Documentation Bot Core

#### 6.1.1 Code Analyzer

```go
// Responsibilities:
// - Parse Go source files using go/parser
// - Extract function signatures, types, interfaces
// - Identify existing documentation
// - Analyze code complexity
// - Extract dependencies and relationships

type Analyzer struct {
    fileSet *token.FileSet
    parser  *parser.Parser
}

type AnalysisResult struct {
    Functions    []FunctionInfo
    Types        []TypeInfo
    Interfaces   []InterfaceInfo
    ExistingDocs map[string]string
    Complexity   ComplexityMetrics
}
```

#### 6.1.2 Claude API Integration

```go
// Responsibilities:
// - Format code context for Claude
// - Build documentation prompts
// - Call Claude API with rate limiting
// - Parse and validate responses
// - Handle errors and retries

type ClaudeClient struct {
    apiKey      string
    model       string
    maxTokens   int
    temperature float64
}

type DocRequest struct {
    CodeContext   string
    ExistingDoc   string
    DocType       string // "godoc", "swagger", "readme"
    Constraints   []string
}

type DocResponse struct {
    Documentation string
    Suggestions   []string
    Confidence    float64
}
```

#### 6.1.3 Documentation Generator

```go
// Responsibilities:
// - Generate godoc-compliant comments
// - Create Swagger annotations
// - Build markdown documentation
// - Generate usage examples
// - Format output correctly

type Generator struct {
    client    *ClaudeClient
    templates map[string]*template.Template
}

// Generates documentation in multiple formats
func (g *Generator) Generate(analysis AnalysisResult) (*Documentation, error)
```

### 6.2 Validation Pipeline

#### 6.2.1 Vale Integration

```yaml
# .vale.ini
StylesPath = .vale/styles
MinAlertLevel = suggestion

[*.md]
BasedOnStyles = Vale, Google, Microsoft
```

#### 6.2.2 Markdownlint Rules

```json
{
  "default": true,
  "MD013": { "line_length": 120 },
  "MD033": false,
  "MD041": false
}
```

#### 6.2.3 Doc Detective Tests

```json
{
  "tests": [
    {
      "id": "api-health-check",
      "description": "Test health check endpoint documentation",
      "steps": [
        {
          "action": "httpRequest",
          "url": "http://localhost:8080/api/v1/health",
          "method": "GET",
          "statusCodes": [200]
        }
      ]
    }
  ]
}
```

---

## 7. GitHub Actions Workflow (Using Official Claude Code Action)

### 7.1 Documentation Generation & PR Creation Workflow

**File:** `.github/workflows/docs-bot.yml`

```yaml
name: Documentation Bot

on:
  pull_request:
    types: [closed]
    branches: [main]

jobs:
  update-docs:
    # Only run if PR was actually merged (not just closed)
    if: github.event.pull_request.merged == true
    runs-on: ubuntu-latest
    permissions:
      contents: write
      pull-requests: write

    steps:
      - name: Checkout repository
        uses: actions/checkout@v4
        with:
          fetch-depth: 0  # Get full history for analysis

      - name: Get merged PR information
        id: pr_info
        run: |
          echo "pr_number=${{ github.event.pull_request.number }}" >> $GITHUB_OUTPUT
          echo "pr_title=${{ github.event.pull_request.title }}" >> $GITHUB_OUTPUT
          echo "pr_author=${{ github.event.pull_request.user.login }}" >> $GITHUB_OUTPUT

          # Fetch PR head commit (critical - it may not be in local main)
          git fetch origin ${{ github.event.pull_request.head.sha }}
          git fetch origin ${{ github.event.pull_request.base.sha }}

          # Get list of changed files
          git diff --name-only ${{ github.event.pull_request.base.sha }} ${{ github.event.pull_request.head.sha }} > /tmp/changed_files.txt
          echo "Changed files:"
          cat /tmp/changed_files.txt

      - name: Run Claude Code to Analyze & Generate Documentation
        uses: anthropics/claude-code-action@v1
        with:
          anthropic_api_key: ${{ secrets.ANTHROPIC_API_KEY }}
          github_token: ${{ secrets.GITHUB_TOKEN }}
          prompt: |
            # Documentation Update Task

            A pull request was just merged into main:

            **PR #${{ steps.pr_info.outputs.pr_number }}:** ${{ steps.pr_info.outputs.pr_title }}
            **Author:** @${{ steps.pr_info.outputs.pr_author }}

            ## Your Task

            1. **Analyze the merged PR changes:**
               - Read the list of changed files from `/tmp/changed_files.txt`
               - Read the actual content of ALL changed files
               - Identify ALL related modules (not just changed files - analyze dependencies and related code)

            2. **Check existing documentation:**
               - Look in the `docs/` folder for existing documentation
               - Check `README.md` for relevant sections
               - Determine if documentation exists for these changes

            3. **Make a decision:**
               - **If documentation is up-to-date:** Create a file `/tmp/no_docs_needed.txt` with message "Documentation is current"
               - **If documentation is missing:** Generate comprehensive new documentation
               - **If documentation is outdated:** Update it with improvements

            4. **Generate/Update documentation (if needed):**

               Create **VERBOSE, COMPREHENSIVE** markdown files in the `docs/` folder:

               For **Backend/API changes:**
               - `docs/API.md` - Complete API reference with all endpoints
               - `docs/ARCHITECTURE.md` - System architecture
               - `docs/DATA_MODELS.md` - Data models and schemas

               For **Frontend changes:**
               - `docs/COMPONENTS.md` - All React components with props/usage
               - `docs/STATE_MANAGEMENT.md` - State management patterns
               - `docs/ROUTING.md` - Routes and navigation

               For **Infrastructure changes:**
               - `docs/INFRASTRUCTURE.md` - All resources explained
               - `docs/DEPLOYMENT.md` - Deployment procedures

               **Always update:**
               - `README.md` - Keep it fresh with latest info

            5. **Documentation style requirements:**
               - Be VERBOSE and DETAILED (not minimal)
               - Include complete code examples (not snippets)
               - Add tables for structured data (props, fields, params)
               - Use collapsible sections for long examples
               - Include mermaid diagrams where appropriate
               - Add cross-references to related docs
               - Include "Last updated" timestamp

            ## Important Notes

            - Analyze ALL related modules, not just changed files
            - If no changes needed, create `/tmp/no_docs_needed.txt` and stop
            - Make documentation self-contained and comprehensive
            - Use examples from the actual codebase

      - name: Install markdownlint
        run: npm install -g markdownlint-cli

      - name: Validate generated documentation
        continue-on-error: true
        run: |
          # Basic formatting check (only if docs exist)
          if compgen -G "docs/**/*.md" > /dev/null 2>&1; then
            markdownlint --fix docs/**/*.md || true
          fi

          # Fix README if it exists
          if [ -f README.md ]; then
            markdownlint --fix README.md || true
          fi

      - name: Check if documentation was updated
        id: check_changes
        run: |
          if [ -f /tmp/no_docs_needed.txt ]; then
            echo "needs_pr=false" >> $GITHUB_OUTPUT
            echo "✅ Documentation is up-to-date. No PR needed."
            cat /tmp/no_docs_needed.txt
          elif git status --short | grep -qE "(docs/|README.md)"; then
            echo "needs_pr=true" >> $GITHUB_OUTPUT
            echo "📝 Documentation changes detected."
            echo ""
            echo "Modified files:"
            git status --short | grep -E "(docs/|README.md)"
          else
            echo "needs_pr=false" >> $GITHUB_OUTPUT
            echo "ℹ️ No documentation changes generated."
          fi

      - name: Create Documentation PR
        if: steps.check_changes.outputs.needs_pr == 'true'
        id: create_pr
        uses: peter-evans/create-pull-request@v6
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
          commit-message: |
            docs: Update documentation for PR #${{ steps.pr_info.outputs.pr_number }}

            Auto-generated comprehensive documentation based on:
            ${{ steps.pr_info.outputs.pr_title }}

            Merged by: @${{ steps.pr_info.outputs.pr_author }}

            🤖 Generated with Claude Code
          branch: docs/update-pr-${{ steps.pr_info.outputs.pr_number }}
          delete-branch: true
          title: "📚 docs: Update for PR #${{ steps.pr_info.outputs.pr_number }}"
          body: |
            ## 📝 Documentation Update

            This PR contains documentation updates based on changes merged in **PR #${{ steps.pr_info.outputs.pr_number }}**.

            ### Original PR

            - **Title:** ${{ steps.pr_info.outputs.pr_title }}
            - **Author:** @${{ steps.pr_info.outputs.pr_author }}
            - **Link:** #${{ steps.pr_info.outputs.pr_number }}

            ### Changes Included

            Documentation has been generated/updated based on analysis of all related modules:

            ```
            $(cat /tmp/changed_files.txt)
            ```

            ### Files Modified

            ```
            $(git status --short | grep -E "(docs/|README.md)")
            ```

            ---

            ### 📋 Review Checklist for Docs Team

            - [ ] **Technical Accuracy** - Verify all technical details are correct
            - [ ] **Code Examples** - Test that code examples actually work
            - [ ] **Completeness** - All relevant aspects are documented
            - [ ] **Links** - All internal and external links work
            - [ ] **Formatting** - Markdown renders correctly
            - [ ] **Clarity** - Documentation is easy to understand
            - [ ] **No Sensitive Data** - No API keys, secrets, or internal URLs

            ---

            🤖 **Auto-generated by Claude Code**

            /cc @docs-team
          labels: |
            documentation
            needs-review
            auto-generated
          # Note: team-reviewers only works in organization repos with a visible team
          # For personal repos, remove this line or use 'reviewers' instead
          # team-reviewers: |
          #   docs-team
          # For personal repos, uncomment this instead:
          # reviewers: |
          #   your-username
          draft: false

      - name: Add comment to PR
        if: steps.create_pr.outputs.pull-request-number != ''
        uses: actions/github-script@v7
        with:
          script: |
            await github.rest.issues.createComment({
              owner: context.repo.owner,
              repo: context.repo.repo,
              issue_number: ${{ steps.create_pr.outputs.pull-request-number }},
              body: `## 🎉 Documentation PR Created!

              **What's in this PR:**
              - Comprehensive markdown documentation
              - Generated by analyzing PR #${{ steps.pr_info.outputs.pr_number }} and all related modules
              - Verbose, detailed content with examples

              **Next Steps:**
              1. Review the generated documentation
              2. Test any code examples
              3. Approve and merge when satisfied

              **Feedback:** If you find issues, comment on specific lines in the "Files changed" tab.`
            });

      - name: Summary
        if: always()
        run: |
          echo "## Documentation Bot Summary" >> $GITHUB_STEP_SUMMARY
          echo "" >> $GITHUB_STEP_SUMMARY
          echo "- **Triggered by:** PR #${{ steps.pr_info.outputs.pr_number }}" >> $GITHUB_STEP_SUMMARY
          echo "- **Documentation needed:** ${{ steps.check_changes.outputs.needs_pr }}" >> $GITHUB_STEP_SUMMARY

          if [ "${{ steps.check_changes.outputs.needs_pr }}" == "true" ]; then
            echo "- **Documentation PR:** #${{ steps.create_pr.outputs.pull-request-number }}" >> $GITHUB_STEP_SUMMARY
            echo "- **PR URL:** ${{ steps.create_pr.outputs.pull-request-url }}" >> $GITHUB_STEP_SUMMARY
          else
            echo "- **Action:** No PR created (docs are current)" >> $GITHUB_STEP_SUMMARY
          fi
```

### 7.2 Manual Trigger Example

You can manually trigger documentation generation:

```bash
gh workflow run "Documentation Bot - Generate & Create PR" \
  -f doc_type=api
```

### 7.3 Workflow Triggers Explained

| Trigger | When | Purpose |
|---------|------|---------|
| `pull_request: [closed]` with `merged == true` | After PR merge to main | Auto-update docs when code changes are merged |
| `workflow_dispatch` | Manual trigger | Generate specific docs on-demand |
| Scheduled (optional) | Weekly cron | Keep docs fresh even without code changes |

---

## 8. Testing Strategy

### 8.1 Test Pyramid

```
                    ┌─────────────┐
                    │   Manual    │
                    │  Review     │
                    └──────┬──────┘
                 ┌─────────┴─────────┐
                 │   E2E Doc Tests   │
                 │  (Doc Detective)  │
                 └─────────┬─────────┘
          ┌──────────────────┴──────────────────┐
          │    Integration Tests               │
          │  (API + Docs validation)           │
          └──────────────────┬──────────────────┘
   ┌────────────────────────────┴────────────────────────────┐
   │              Unit Tests                                  │
   │  - Doc generation tests                                  │
   │  - Parser tests                                          │
   │  - Validator tests                                       │
   └──────────────────────────────────────────────────────────┘
```

### 8.2 Test Categories

#### 8.2.1 Unit Tests

```go
// Test documentation generation for each component
func TestGenerateGodocComment(t *testing.T)
func TestGenerateSwaggerAnnotation(t *testing.T)
func TestParseGoFile(t *testing.T)
func TestValidateDocumentation(t *testing.T)
```

#### 8.2.2 Integration Tests

```go
// Test end-to-end documentation workflow
func TestDocumentationWorkflow(t *testing.T)
func TestSwaggerGeneration(t *testing.T)
func TestClaudeAPIIntegration(t *testing.T)
```

#### 8.2.3 Quality Tests

- Vale prose linting
- Markdownlint formatting
- Link validation
- Code example compilation
- API endpoint verification

#### 8.2.4 Accuracy Tests

- Doc Detective scenarios
- API response validation
- Example code execution
- Documentation-code consistency

---

## 9. Documentation Quality Metrics

### 9.1 Measurable Metrics

```
┌──────────────────────────────────────────────────────┐
│         Documentation Quality Scorecard              │
├──────────────────────────────────────────────────────┤
│                                                      │
│  Coverage Metrics:                                   │
│  ├─ Functions documented: X/Y (Z%)                  │
│  ├─ Types documented: X/Y (Z%)                      │
│  ├─ Public APIs documented: X/Y (Z%)                │
│  └─ Examples provided: X/Y (Z%)                     │
│                                                      │
│  Quality Metrics:                                    │
│  ├─ Vale score: X/100                               │
│  ├─ Markdownlint issues: X                          │
│  ├─ Broken links: X                                 │
│  └─ Godoc warnings: X                               │
│                                                      │
│  Accuracy Metrics:                                   │
│  ├─ Doc Detective tests passed: X/Y (Z%)           │
│  ├─ API examples verified: X/Y (Z%)                │
│  └─ Code examples compile: X/Y (Z%)                │
│                                                      │
│  Overall Score: X/100                                │
└──────────────────────────────────────────────────────┘
```

### 9.2 Quality Gates

| Metric | Minimum Threshold | Target |
|--------|------------------|---------|
| Function documentation coverage | 80% | 95% |
| Public API documentation | 100% | 100% |
| Vale score | 70 | 90 |
| Markdownlint issues | 10 | 0 |
| Broken links | 0 | 0 |
| Doc Detective pass rate | 90% | 100% |
| Example compilation rate | 100% | 100% |

---

## 10. Prompt Engineering for Verbose Markdown Documentation

### 10.1 API Documentation Generation Prompt

```
You are an expert technical writer creating comprehensive API documentation.

Task: Generate VERBOSE, DETAILED documentation for the following API endpoint in markdown format.

Code Context:
{handler_code}

Models:
{model_definitions}

Existing Documentation (if any):
{existing_doc}

Requirements:
1. **Comprehensive Coverage**: Include ALL the following sections:
   - Endpoint description (2-3 paragraphs explaining purpose and use cases)
   - Authentication requirements
   - Request format (headers, body, query params)
   - Request schema table (field, type, required, constraints, description)
   - Response format for success
   - Response schema table
   - ALL possible error responses with status codes
   - Code examples in 4+ languages (cURL, Go, Python, JavaScript)
   - Business logic notes
   - Rate limiting information
   - Related endpoints with links

2. **Verbosity**: Be DETAILED, not minimal:
   - Explain WHY things work the way they do
   - Provide context about business logic
   - Include helpful tips and gotchas
   - Show realistic, production-ready examples

3. **Code Examples**: Make them complete and runnable:
   - Include error handling
   - Show proper imports
   - Use realistic variable names
   - Add comments explaining key parts

4. **Formatting**: Use proper markdown:
   - Tables for structured data
   - Code blocks with language tags
   - Collapsible sections for examples
   - Proper headers (###, ####)

Generate the complete API documentation in markdown format:
```

### 10.2 Architecture Documentation Prompt

```
You are a senior software architect creating comprehensive architecture documentation.

Task: Generate VERBOSE, DETAILED system architecture documentation in markdown format.

Codebase Analysis:
- Project structure: {project_structure}
- Dependencies: {dependencies}
- Key components: {components}
- Data flows: {data_flows}

Existing Documentation (if any):
{existing_doc}

Requirements:
1. **Comprehensive Sections**:
   - Overview (explain the "why" behind architectural decisions)
   - High-level architecture diagram (ASCII art)
   - Component details (each component gets its own subsection)
   - Data flow diagrams
   - Technology stack with justifications
   - Design decisions and trade-offs
   - Security considerations
   - Performance characteristics
   - Deployment architecture
   - Scaling considerations

2. **Depth**: Go DEEP into each topic:
   - Explain patterns used and why
   - Document responsibilities of each layer
   - Show component interactions
   - Discuss alternatives considered

3. **Diagrams**: Include ASCII diagrams:
   - System architecture
   - Component relationships
   - Data flows
   - Deployment topology

4. **Length**: Aim for 3000+ words with detailed explanations

Generate the complete architecture documentation in markdown format:
```

### 10.3 README Enhancement Prompt

```
You are a developer advocate creating an exceptional README for a GitHub project.

Task: Generate or IMPROVE the README.md with COMPREHENSIVE, DETAILED content.

Project Information:
- Code structure: {project_structure}
- Main features: {features}
- Dependencies: {dependencies}
- API endpoints: {endpoints}

Existing README (if any):
{existing_readme}

Requirements:
1. **Essential Sections** (all required):
   - Project title and tagline
   - Badges (build status, coverage, license, etc.)
   - Table of contents (with anchor links)
   - Overview (2-3 paragraphs)
   - Features (detailed list with explanations)
   - Getting Started
     * Prerequisites (with version requirements and install instructions)
     * Installation (step-by-step with commands)
     * Configuration (environment variables, config files)
   - Usage
     * Running locally (detailed steps)
     * API endpoints (summary table)
     * Example requests
   - Documentation (links to other docs)
   - Development
     * Project structure explanation
     * Building (with commands)
     * Testing (with commands)
     * Code standards
   - Deployment (overview)
   - Contributing (guidelines)
   - License

2. **Verbosity**: Make it COMPREHENSIVE:
   - Assume reader knows nothing about the project
   - Provide copy-paste ready commands
   - Explain what each step does
   - Add troubleshooting tips
   - Include system requirements

3. **Examples**: Show real, working examples:
   - Actual API calls with responses
   - Configuration file examples
   - Common use cases

4. **Visual Appeal**:
   - Use emojis for section headers
   - Include badges
   - Use tables for structured data
   - Add collapsible sections for optional content

Generate the complete, enhanced README.md in markdown format:
```

### 10.4 Usage/Tutorial Documentation Prompt

```
You are a technical educator creating comprehensive usage documentation and tutorials.

Task: Generate DETAILED, TUTORIAL-STYLE documentation showing how to use this system.

System Information:
- Features: {features}
- API endpoints: {endpoints}
- Use cases: {use_cases}

Existing Documentation (if any):
{existing_doc}

Requirements:
1. **Tutorial Structure**:
   - Introduction (what will be learned)
   - Prerequisites
   - Step-by-step walkthrough
   - Common patterns
   - Advanced usage
   - Troubleshooting
   - FAQ

2. **Teaching Style**:
   - Explain concepts before showing code
   - Start simple, build complexity gradually
   - Show output after each step
   - Explain WHY, not just HOW
   - Include "what if" scenarios

3. **Code Examples**:
   - Complete, runnable examples
   - Progressive examples (basic → intermediate → advanced)
   - Real-world scenarios
   - Error handling patterns

4. **Completeness**:
   - Cover all major features
   - Include edge cases
   - Show best practices
   - Warn about common pitfalls

Generate the complete usage/tutorial documentation in markdown format:
```

---

## 11. Implementation Phases (REVISED - MVP First)

### ⚠️ CRITICAL: Scope Reduction for MVP

Based on architectural review, the original phases were too ambitious. **Start small and prove value quickly.**

### Phase 1: MVP - Auto-Generate Docs via PRs (Week 1-2) 🎯

**Goal**: Demonstrate value by automatically creating comprehensive documentation PRs

**Scope**:
- Create simple Go API with 3-5 endpoints (health, users, products)
- Use official Claude Code GitHub Action (not custom bot)
- Trigger on merged PRs (not during PR review)
- Analyze ALL related Go modules (not just changed files)
- Generate **verbose markdown documentation files**:
  - `docs/API.md` - Complete API reference
  - `docs/ARCHITECTURE.md` - System architecture
  - `docs/DATA_MODELS.md` - Go structs
  - `README.md` updates
- **Create new PR** with documentation file changes (not inline comments)
- Assign PR to docs-team for review
- Smart detection: Only create PR if docs actually need updating

**Deliverables**:
```
✓ Simple Go API running (cmd/api/main.go)
✓ GitHub Actions workflow using Claude Code Action
✓ Triggers on pull_request: [closed] + merged == true
✓ Git diff analysis with proper fetch of PR commits
✓ Safe markdownlint with conditional file checking
✓ Creates documentation PR only when needed
✓ Docs team review workflow
```

**Success Metrics**:
- Bot runs successfully after PR merge
- Generates comprehensive documentation files
- Only creates PR when docs are missing/outdated
- Docs team can review and approve before merging
- No false positives (creating PRs when docs are current)

### Phase 2: Add Basic Validation (Week 3)

**Scope**:
- Add godoc validation (format checking)
- Integrate markdownlint for README only
- Add simple metrics (coverage %)
- Cache tool installations in Actions

**Skip for now**: Vale, Doc Detective, Swagger generation

### Phase 3: Swagger for Changed API Routes (Week 4-5)

**Scope**:
- Detect changed HTTP handlers
- Generate Swagger annotations ONLY for changed endpoints
- Add swaggo/swag integration
- Gate behind label: `docs:swagger` or when `internal/handlers/` changes

**Heuristic**:
```go
if pr.HasLabel("docs:swagger") || pr.ChangesMatch("internal/handlers/*.go") {
    runSwaggerGeneration()
}
```

### Phase 4: Validation & Feedback Loop (Week 6-7)

**Scope**:
- Add Vale for markdown (cached install)
- Add reaction mechanism: 👍/👎 on suggestions
- Store feedback in GitHub Discussions or issue comments
- Use feedback to tune prompts

### Phase 5: Expansion & Optimization (Week 8+)

**Scope**:
- Add Doc Detective (only when `docs:test` label present)
- Implement comprehensive caching
- Add diff-based summarization
- Optimize prompt templates based on feedback

---

## 12. Cost Optimization Strategy

### 12.1 Claude API Cost Management

**Critical for Production Use**

#### Token Limits

```go
const (
    MaxTokensPerFunction = 2000    // ~$0.006 per function
    MaxTokensPerRequest  = 8000    // ~$0.024 per request
    MaxFunctionsPerPR    = 20      // ~$0.12 per PR
    MaxCostPerPR         = 0.50    // Hard limit: $0.50
)

type CostTracker struct {
    inputTokens  int
    outputTokens int
    requestCount int
}

func (ct *CostTracker) EstimatedCost() float64 {
    // Claude Sonnet 4.5 pricing (as of 2025)
    inputCost := float64(ct.inputTokens) * 0.000003   // $3/M tokens
    outputCost := float64(ct.outputTokens) * 0.000015 // $15/M tokens
    return inputCost + outputCost
}

func (ct *CostTracker) ExceedsLimit() bool {
    return ct.EstimatedCost() > MaxCostPerPR
}
```

#### Rate Limiting

```go
type RateLimiter struct {
    requestsPerMinute int
    concurrency       int
    backoff           time.Duration
}

// Start with conservative limits
func NewRateLimiter() *RateLimiter {
    return &RateLimiter{
        requestsPerMinute: 10,   // Stay well under API limits
        concurrency:       2,    // Process 2 functions in parallel max
        backoff:           time.Second * 5,
    }
}
```

#### Caching Strategy

```go
// Cache generated documentation by function signature hash
type DocCache struct {
    store map[string]CachedDoc // key: SHA256(funcSignature + codeBody)
    ttl   time.Duration         // 7 days
}

// Only call Claude API if:
// 1. Function signature changed
// 2. Function body changed
// 3. Cache expired
// 4. Existing doc is missing
```

### 12.2 GitHub Actions Cost Management

**Free Tier Limits**: 2000 minutes/month for private repos

#### Optimization Strategies (MVP Approach)

1. **Conditional Execution**:
```yaml
jobs:
  update-docs:
    runs-on: ubuntu-latest
    # Only run when PR is actually merged, not just closed
    if: github.event.pull_request.merged == true
```

2. **Smart Detection**:
```yaml
# Bot creates /tmp/no_docs_needed.txt if docs are current
# Workflow checks this file before creating PR
- name: Check if documentation was updated
  id: check_changes
  run: |
    if [ -f /tmp/no_docs_needed.txt ]; then
      echo "needs_pr=false" >> $GITHUB_OUTPUT
    fi
```

3. **Tool Caching**:
```yaml
- name: Cache Vale
  uses: actions/cache@v3
  with:
    path: ~/.vale
    key: vale-${{ hashFiles('.vale.ini') }}

- name: Cache Go modules
  uses: actions/cache@v3
  with:
    path: ~/go/pkg/mod
    key: go-${{ hashFiles('go.sum') }}
```

### 12.3 Cost Monitoring

```yaml
- name: Report Costs
  run: |
    echo "📊 Documentation Bot Cost Report" >> $GITHUB_STEP_SUMMARY
    echo "- Functions analyzed: $FUNC_COUNT" >> $GITHUB_STEP_SUMMARY
    echo "- Claude API calls: $API_CALLS" >> $GITHUB_STEP_SUMMARY
    echo "- Input tokens: $INPUT_TOKENS" >> $GITHUB_STEP_SUMMARY
    echo "- Output tokens: $OUTPUT_TOKENS" >> $GITHUB_STEP_SUMMARY
    echo "- Estimated cost: \$$ESTIMATED_COST" >> $GITHUB_STEP_SUMMARY
    echo "- Actions runtime: $RUNTIME minutes" >> $GITHUB_STEP_SUMMARY
```

---

## 13. Execution Heuristics - When to Run What

**⚠️ NOTE: This section describes FUTURE enhancements (Phase 3+). Current MVP has simpler logic.**

**Current MVP (Phase 1) Decision Tree:**

```
PR Merged to Main
│
├─ Was PR actually merged?
│  └─ NO → Skip (exit early)
│  └─ YES → Continue
│
├─ Get changed files from PR
│
├─ Analyze ALL related Go modules
│  ├─ Read changed files
│  ├─ Parse dependencies
│  └─ Map related modules
│
├─ Check existing documentation
│  └─ docs/ folder
│  └─ README.md
│
├─ Documentation current?
│  └─ YES → Create /tmp/no_docs_needed.txt (no PR)
│  └─ NO → Generate comprehensive markdown docs
│
└─ Documentation generated?
   └─ YES → Create new PR for docs team
   └─ NO → End (no changes)
```

**Future Enhancement (Phase 3+) Decision Tree:**

Original complex decision tree with labels, file triggers, etc. will be implemented in later phases.

### 13.2 Label-Based Gating (Future Phase)

**⚠️ NOTE: MVP does not use labels. This is for Phase 3+ enhancement.**

| Label | Triggers | Reason |
|-------|----------|--------|
| `docs:swagger` | Swagger generation + validation | Slow (30s-2min), only needed for API changes |
| `docs:validate` | Vale + markdownlint | Medium speed (10-30s), run on demand |
| `docs:test` | Doc Detective accuracy tests | Very slow (1-5min), requires app startup |
| `docs:skip` | Skip all documentation checks | Emergency escape hatch |
| `docs:full` | Run everything | For release candidates |

### 13.3 File Path Triggers (Future Phase)

**⚠️ NOTE: MVP analyzes all related modules. This is for Phase 3+ enhancement.**

```go
type TriggerRules struct {
    SwaggerPaths []string
    SkipPaths    []string
}

var DefaultRules = TriggerRules{
    SwaggerPaths: []string{
        "internal/handlers/**/*.go",
        "pkg/api/**/*.go",
    },
    SkipPaths: []string{
        "**/*_test.go",
        "vendor/**",
        "docs/**",
        ".github/**",
    },
}

func ShouldRunSwagger(changedFiles []string) bool {
    for _, file := range changedFiles {
        if matches(file, DefaultRules.SwaggerPaths) &&
           !matches(file, DefaultRules.SkipPaths) {
            return true
        }
    }
    return false
}
```

### 13.4 Size-Based Throttling

```go
// For large PRs, limit scope to prevent cost explosion
func SelectFunctionsToDocument(functions []Function) []Function {
    const maxFunctions = 20

    if len(functions) <= maxFunctions {
        return functions
    }

    // Prioritize:
    // 1. Exported functions without docs
    // 2. Exported functions with poor docs
    // 3. Complex functions (high cyclomatic complexity)

    sorted := prioritizeFunctions(functions)
    return sorted[:maxFunctions]
}
```

---

## 14. Performance Optimization

### 14.1 Parallel Processing

```go
// Process multiple functions concurrently
func (g *Generator) GenerateDocs(functions []Function) ([]Doc, error) {
    const maxConcurrency = 3 // Conservative to respect rate limits

    semaphore := make(chan struct{}, maxConcurrency)
    results := make(chan Result, len(functions))

    for _, fn := range functions {
        go func(f Function) {
            semaphore <- struct{}{} // Acquire
            defer func() { <-semaphore }() // Release

            doc, err := g.generateForFunction(f)
            results <- Result{doc, err}
        }(fn)
    }

    // Collect results...
}
```

### 14.2 Incremental Analysis

```go
// Only analyze changed files, not entire codebase
type DiffAnalyzer struct {
    baseRef string
    headRef string
}

func (da *DiffAnalyzer) GetChangedFunctions() ([]Function, error) {
    // git diff --name-only base..head
    changedFiles := da.getChangedFiles()

    var functions []Function
    for _, file := range changedFiles {
        if !strings.HasSuffix(file, ".go") {
            continue
        }

        // Parse only changed files
        funcs := parseFile(file)
        functions = append(functions, funcs...)
    }

    return functions, nil
}
```

### 14.3 Prompt Optimization

```go
// Keep prompts concise to reduce token usage
type OptimizedPrompt struct {
    MaxContextLines  int // 50 lines max around function
    IncludeImports   bool // false - not needed for docs
    IncludeDeps      bool // false - focus on function only
}

func BuildPrompt(fn Function) string {
    // Include only:
    // 1. Function signature
    // 2. 10 lines of body (not full implementation)
    // 3. Existing doc (if any)
    // 4. Parameter types (not full definitions)

    return fmt.Sprintf(`Generate godoc for:

%s

Existing doc: %s
Keep it under 5 sentences.`,
        fn.SignatureWithPartialBody(10),
        fn.ExistingDoc,
    )
}
```

### 14.4 GitHub Actions Optimization

```yaml
jobs:
  doc-bot:
    runs-on: ubuntu-latest
    timeout-minutes: 10  # Hard limit

    steps:
      # Use composite actions to avoid repeating setup
      - uses: ./.github/actions/setup-go-env

      # Cache everything possible
      - uses: actions/cache@v3
        with:
          path: |
            ~/.cache/go-build
            ~/go/pkg/mod
            ~/.vale
            node_modules
          key: ${{ runner.os }}-docbot-${{ hashFiles('**/go.sum', '.vale.ini', 'package-lock.json') }}

      # Run steps in parallel where possible
      - name: Parallel validation
        run: |
          vale docs/ & \
          markdownlint '**/*.md' & \
          wait
```

---

## 15. PR UX Strategy

### 15.1 Suggestions vs Auto-Commits

**Decision: Use Suggested Changes (GitHub's native feature)**

#### Why Suggested Changes?

✅ **Pros**:
- Reviewers can accept/reject with one click
- Non-intrusive (doesn't pollute commit history)
- Easy to edit before accepting
- Native GitHub UI support
- Doesn't trigger additional CI runs
- Clear attribution (bot vs human)

❌ **Auto-Commit Cons**:
- Frustrates reviewers who want to tweak
- Requires additional push (triggers CI again)
- Makes commit history noisy
- Harder to attribute changes
- May conflict with human edits

---

**⚠️ NOTE: This section describes a FUTURE alternative approach (Phase 3+), not the current MVP.**

**Current MVP (Phase 1)**: Bot creates new PRs with markdown documentation files.

**Future Option (Phase 3+)**: Bot could optionally post inline suggestions as PR review comments.

---

#### Implementation (Future Phase)

```go
// Post as GitHub review with suggested changes
func (pr *PRClient) PostSuggestion(file string, line int, suggestion string) error {
    comment := fmt.Sprintf(`## 📝 Documentation Suggestion

%s

<details>
<summary>Why this suggestion?</summary>

- Missing godoc comment for exported function
- Follows Go documentation conventions
- Generated by Claude Sonnet 4.5

React with 👍 if helpful, 👎 if not relevant.
</details>`, suggestion)

    return pr.CreateReviewComment(&github.ReviewComment{
        Path:      file,
        Line:      line,
        Body:      comment,
        StartLine: line - 1,
        Side:      "RIGHT",
    })
}
```

### 15.2 PR Comment Format

```markdown
## 🤖 Documentation Bot Report

✅ **Analysis complete** - Reviewed 5 changed functions

### 📊 Summary
- Functions analyzed: 5
- Suggestions posted: 3
- Already documented: 2
- Cost: $0.08

### 📝 Suggestions Posted
1. `CreateUser` - Missing godoc comment
2. `GetUserByID` - Incomplete parameter documentation
3. `DeleteUser` - Missing error documentation

### 💡 Next Steps
- Review suggestions inline
- Accept/reject with GitHub's suggestion UI
- React with 👍/👎 to help improve suggestions

---
<sub>Powered by Claude Code | Took 1m 23s | [Report Issue](link)</sub>
```

### 15.3 Inline vs Summary

**Strategy**: Hybrid approach

1. **Inline comments**: For specific function suggestions
2. **Summary comment**: For overall PR analysis
3. **No notification spam**: Batch all comments into one review

```go
// Batch comments into single review
review := &github.Review{
    Event: "COMMENT", // Not REQUEST_CHANGES
    Body:  summaryComment,
    Comments: []github.ReviewComment{
        // All inline suggestions
    },
}
```

---

## 16. Feedback Loop Mechanisms

**⚠️ NOTE: This section describes FUTURE enhancements (Phase 3+), not the current MVP.**

**Current MVP (Phase 1)**: Bot creates documentation PRs; docs team reviews and approves/rejects.

**Future Enhancement (Phase 3+)**: Could add feedback tracking for continuous improvement.

---

### 16.1 Reaction-Based Feedback (Future Phase)

```go
// Track reactions on bot comments
type FeedbackTracker struct {
    db Database
}

func (ft *FeedbackTracker) RecordReaction(commentID int, reaction string) {
    // Store: comment_id, reaction (+1/-1), function_signature, prompt_version
    ft.db.Insert(Feedback{
        CommentID:    commentID,
        Reaction:     reaction, // 👍 or 👎
        FunctionSig:  extractFunctionSig(commentID),
        PromptVersion: "v1.2.3",
        Timestamp:    time.Now(),
    })
}
```

### 16.2 Feedback Analysis

```sql
-- Query to find low-quality suggestions
SELECT function_pattern, COUNT(*) as negative_reactions
FROM feedback
WHERE reaction = '👎'
GROUP BY function_pattern
HAVING COUNT(*) > 5
ORDER BY negative_reactions DESC;

-- Find what works well
SELECT prompt_version, COUNT(*) as positive_reactions
FROM feedback
WHERE reaction = '👍'
GROUP BY prompt_version
ORDER BY positive_reactions DESC;
```

### 16.3 Prompt Tuning Based on Feedback

```go
// Version prompts and A/B test
type PromptRegistry struct {
    versions map[string]PromptTemplate
    router   *ABTestRouter
}

func (pr *PromptRegistry) GetPrompt(fnType string) PromptTemplate {
    // Route 10% of traffic to new prompt version
    if pr.router.ShouldUseExperimental() {
        return pr.versions["experimental-v2"]
    }
    return pr.versions["stable-v1"]
}
```

### 16.4 Manual Override System

```yaml
# .github/docbot.yml - Allow teams to configure behavior
feedback:
  enabled: true
  reactions: ["+1", "-1"]

prompts:
  godoc:
    version: "v1.2.3"
    max_length: 5  # sentences
    include_examples: false

  swagger:
    version: "v1.0.1"
    include_responses: [200, 400, 500]

overrides:
  # Skip specific functions
  skip_functions:
    - "internal/legacy/**"

  # Always document these
  require_docs:
    - "pkg/api/**"
```

### 16.5 Continuous Improvement Dashboard (Future Phase)

**⚠️ NOTE: MVP uses simpler metrics. This dashboard is for Phase 3+.**

**MVP Dashboard (Phase 1):**

```
📊 Documentation Bot Health (Last 30 Days)

┌─────────────────────────────────────────┐
│ PRs Merged:             45              │
│ Documentation PRs Created: 12           │
│ PRs Merged by Docs Team: 10             │
│ PRs Rejected/Closed:    2               │
│ False Positives:        1 (8%)          │
│ Avg Workflow Runtime:   3m 12s          │
└─────────────────────────────────────────┘

Top Issues:
1. Documentation too verbose (2 PRs edited heavily)
2. Missed related module (1 PR)
3. Created PR when docs current (1 false positive)

Documentation Quality:
- PRs approved with no edits: 8 (80%)
- PRs requiring minor edits: 2 (20%)
- PRs rejected as unnecessary: 2
```

---

## 17. Success Criteria (Revised for MVP)

### 17.1 Phase 1 MVP Requirements (Week 1-2)

✓ Bot analyzes merged PRs and ALL related Go modules (not just changed files)
✓ Bot checks if documentation exists and is up-to-date
✓ Bot generates comprehensive verbose markdown files (API.md, ARCHITECTURE.md, DATA_MODELS.md, README.md)
✓ Creates new PR with documentation file changes (not inline comments)
✓ Only creates PR if documentation actually needs updating
✓ Assigns PR to docs-team for review
✓ Workflow completes in reasonable time
✓ Smart detection prevents unnecessary PRs

### 17.2 Phase 2 Requirements (Week 3)

✓ Basic godoc format validation
✓ markdownlint for README only
✓ Coverage metrics displayed
✓ Tool installations cached

### 17.3 Phase 3+ Requirements (Week 4+)

✓ Swagger generation (label-gated)
✓ Vale integration (label-gated)
✓ Feedback loop with 👍/👎 reactions
✓ Doc Detective tests (label-gated)
✓ A/B testing for prompt optimization

---

## 18. Useful Resources

### 18.1 Documentation Tools

- **Vale**: https://vale.sh/
- **markdownlint**: https://github.com/DavidAnson/markdownlint
- **Doc Detective**: https://github.com/doc-detective/doc-detective
- **swaggo/swag**: https://github.com/swaggo/swag
- **llm-docs-builder**: https://github.com/karafka/llm-docs-builder

### 18.2 Go Documentation

- **Go Doc Comments**: https://tip.golang.org/doc/comment
- **Godoc**: https://go.dev/blog/godoc
- **Effective Go**: https://go.dev/doc/effective_go

### 18.3 Claude API

- **Anthropic API Docs**: https://docs.anthropic.com/
- **Claude API Guide**: https://www.anthropic.com/learn/build-with-claude
- **Best Practices**: https://www.anthropic.com/engineering/claude-code-best-practices

### 18.4 GitHub Actions

- **Actions Documentation**: https://docs.github.com/en/actions
- **Workflow Syntax**: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions
- **Create Pull Request Action**: https://github.com/peter-evans/create-pull-request

---

## 19. Alternative Approaches Considered

### 19.1 Local vs Cloud-Based Bot

| Aspect | Local Bot | Cloud-Based Bot |
|--------|-----------|-----------------|
| **Execution** | Runs in GitHub Actions | Runs on dedicated server |
| **Cost** | Free (Actions minutes) | Server + hosting costs |
| **Scalability** | Limited by Actions | Highly scalable |
| **Setup** | Simple | Complex |
| **Security** | API key in secrets | More secure isolation |

**Decision**: Start with GitHub Actions (simpler), migrate to cloud if needed.

### 19.2 Documentation Testing Approaches

| Approach | Pros | Cons |
|----------|------|------|
| **Vale + markdownlint** | Fast, syntax-focused | No accuracy testing |
| **Doc Detective** | Tests accuracy | Requires running app |
| **Custom test framework** | Fully customizable | High maintenance |
| **Hybrid approach** | Comprehensive coverage | More complex setup |

**Decision**: Start with markdownlint only, add Vale when needed (Phase 2+).

### 19.3 AI Model Selection

| Model | Pros | Cons |
|-------|------|------|
| **Claude Sonnet 4.5** | Excellent at code understanding, fast | API costs |
| **GPT-4** | Good documentation | Less code-focused |
| **Local LLM** | No API costs | Lower quality, resource-intensive |
| **Specialized model** | Purpose-built | Limited availability |

**Decision**: Claude Sonnet 4.5 (best for code + documentation tasks).

---

## 20. Risk Mitigation (Updated)

### 20.1 Identified Risks

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| API costs too high | High | Medium | Cache results, batch requests, set limits |
| Generated docs incorrect | High | Low | Multi-layer validation, manual review |
| CI/CD pipeline too slow | Medium | Medium | Parallel execution, incremental analysis |
| API rate limits hit | Medium | High | Implement exponential backoff, queue |
| Claude API downtime | Low | Low | Graceful degradation, manual fallback |

### 20.2 Fallback Strategies (MVP)

1. **API Failure**: Workflow fails gracefully, comment on merged PR with error
2. **Generation Failure**: Create GitHub issue for manual documentation review
3. **Performance Issues**: Workflow timeout (10 minutes max), fail and notify
4. **Quality Issues**: Docs team reviews and provides feedback via PR comments

---

## Conclusion (Revised)

### Key Takeaways

This design document presents a **pragmatic, incremental approach** to building an AI-powered documentation bot for Go applications.

**Critical Success Factors:**

1. **Start Small** - MVP focuses on simple Go API with comprehensive markdown documentation
2. **Prove Value** - Week 1-2 goal: Successfully generate documentation for merged PRs
3. **Smart Detection** - Only create PR when documentation actually needs updating
4. **Comprehensive Analysis** - Analyze ALL related modules, not just changed files
5. **Docs Team Workflow** - Create new PRs for docs team to review, not auto-commits
6. **Extensible Design** - Architecture supports future expansion to other languages

**What Makes This Design Practical:**

✅ **Incremental rollout** - Start with Go API demo, expand to other languages later
✅ **Official tooling** - Uses Claude Code GitHub Action (not custom implementation)
✅ **Comprehensive analysis** - Analyzes all related modules for complete documentation
✅ **Smart detection** - Only creates PR when documentation needs updating
✅ **Non-intrusive** - New PRs for review, not auto-commits; docs team stays in control
✅ **Extensible architecture** - Designed to support React, Python, Terraform, and more

**What Changed from Original Design:**

- ❌ ~~Complex custom bot implementation~~ → ✅ Use official Claude Code GitHub Action
- ❌ ~~Inline PR review comments~~ → ✅ Create new PRs with markdown files
- ❌ ~~Changed files only~~ → ✅ Analyze ALL related modules for completeness
- ❌ ~~Minimal documentation~~ → ✅ Verbose, comprehensive markdown files
- ❌ ~~Always create PR~~ → ✅ Smart detection: only when docs need updating
- ❌ ~~Multi-language from day 1~~ → ✅ Go-first with extensible architecture

---

## Next Steps: Actionable Implementation Plan

### Week 1 - Day 1-2: Project Setup

```bash
# 1. Create basic Go API structure
mkdir -p cmd/api internal/handlers pkg/docbot
go mod init github.com/yourorg/doc-agent-demo

# 2. Implement simple API (3-5 endpoints)
# - Health check
# - User CRUD (create, get, list)
# - Product CRUD (create, get, list)

# 3. Add Anthropic API key to GitHub Secrets
gh secret set ANTHROPIC_API_KEY
```

### Week 1 - Day 3-4: Core Bot Logic (MVP)

**⚠️ NOTE: MVP uses official tooling. Custom implementation is for future phases.**

```bash
# 4. Configure workflow to use official action
# .github/workflows/docs-bot.yml
# ✅ Already complete - using anthropics/claude-code-action@v1

# 5. Write comprehensive prompt for documentation generation
# Prompt engineering in workflow YAML
# ✅ Already complete - analyzes all modules, generates verbose docs

# 6. Test smart detection logic
# Verify /tmp/no_docs_needed.txt mechanism works
```

**Future Enhancement (Phase 3+):**
Custom implementation with diff analyzer, AST parser, and direct API client would be built in later phases if needed.

### Week 1 - Day 5-7: GitHub Integration (MVP)

```bash
# 7. Create GitHub Actions workflow
# .github/workflows/docs-bot.yml
# ✅ Already complete - using Claude Code GitHub Action

# 8. Set up Anthropic API key secret
# gh secret set ANTHROPIC_API_KEY
# ✅ Already complete

# 9. Test by merging sample PRs
# Verify documentation PRs are created correctly
```

### Week 2: Polish & Iterate (MVP)

```bash
# 10. Improve prompt for better documentation quality
# 11. Test smart detection (no PR when docs current)
# 12. Verify docs team workflow
# 13. Test with various PR types
# 14. Gather feedback from docs team
```

### Success Metrics for Week 1-2 (MVP)

- [ ] Bot successfully runs when PRs are merged
- [ ] Documentation PRs created only when needed
- [ ] Comprehensive markdown files generated (API.md, ARCHITECTURE.md, etc.)
- [ ] Zero false positives (PRs created when docs are current)
- [ ] Docs team can review and approve PRs
- [ ] Workflow completes in reasonable time

### When to Move to Phase 2

Move to Phase 2 (validation) when:
- ✅ 10+ documentation PRs created and reviewed
- ✅ Docs team satisfied with quality
- ✅ Smart detection working (no unnecessary PRs)
- ✅ Documentation is accurate and comprehensive
- ✅ Prompt templates stabilized

### Questions to Answer Before Building

1. ❓ What's your Anthropic API budget per month?
2. ❓ How many PRs per week does your team create?
3. ❓ Should bot run on public repos or private only?
4. ❓ Who will review documentation PRs and provide feedback?
5. ❓ What's the rollout plan? (Single repo first? All repos?)

---

## Final Recommendation

**Build Phase 1 MVP first.** The current implementation provides a solid foundation using official tooling. Don't try to implement everything at once.

**Timeline:**
- Week 1-2: MVP (auto-generated docs PRs after merge) ✅ Complete
- Week 3: Test with real PRs, gather docs team feedback
- Week 4-5: Iterate on prompts and documentation quality
- Week 6+: Add validation tools or expand to other languages

**Remember:** A simple bot that generates comprehensive documentation PRs is infinitely more valuable than a complex bot that tries to do everything at once.

Good luck! 🚀
