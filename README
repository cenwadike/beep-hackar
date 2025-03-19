# Beep; Mobile money on blockchain

## Project Overview

Beep is a mobile money platform built on the blockchain, designed to bring seamless financial 
transactions to users by tokenizing the Nigerian Naira (NGN) and integrating with the Cosmos 
ecosystem. Users can deposit Naira to receive Beep tokens, transfer these tokens to others, 
withdraw Naira to their bank accounts, and create intents to exchange Naira for tAtom tokens 
(a tokenized representation of ATOM). Deployed on the Neutron blockchain, Beep leverages the 
scalability and interoperability of Cosmos technologies to enable efficient, secure, and 
transparent mobile money services with tokenized assets.

The source code is available at https://github.com/come-senusi-wale/beep-hackar. Beep demonstrates 
a practical application of blockchain in mobile financial services, bridging traditional fiat with 
tokenized ecosystems.

## Key Features

- **Tokenize Naira**: Deposit Naira and receive equivalent tNGN tokens.
- **Transfer Tokens**: Send Beep tokens to other users seamlessly.
- **Withdraw Naira**: Convert Beep tokens back to Naira and withdraw to a bank account.
- **Swap Tokens**: Exchange Naira for tAtom tokens or vice-versa.

## Technical Architecture

Beep combines a Node.js backend with USSD integration and smart contracts on Neutron, using MongoDB 
for off-chain data storage. Below is its technical structure:

### USSD Interface

- Access Point: USSD service (e.g., *123#), enabling users to interact via mobile phones without 
internet access.
- Implementation: Integrated with a USSD gateway (e.g., Termii, inferred from TERMII_API_KEY), processing 
menu-driven commands.
- Functionality:
    - Deposit Naira and mint Beep tokens.
    - Transfer tokens by entering recipient details.
    - Request Naira withdrawals.
    - Exchange tNGN to tATOM.

### Backend

- Framework: Node.js with Express.js, handling API endpoints for user interactions.
- Database: MongoDB, storing user data, transaction records, and intents (off-chain).
- Authentication: JWT (JSON Web Tokens) for secure user sessions.
- API Endpoints:
    - Tokenization of Naira to Beep tokens.
    - Token transfers between users.
    - Naira withdrawal requests.
    - Intent creation for token exchange.

### Blockchain Layer

- Platform: Neutron blockchain, hosting CosmWasm smart contracts.
- Contracts:
    - Beep Token Contract: Manages Beep token minting, transfers, and burning.
    - Mono-Chain-Beep Contract: Facilitates intents for token exchange.
- Integration: Cosmos SDK and cosmos.js for blockchain interactions from the backend.

## Data Flow

- User deposits Naira via the frontend, triggering a backend API call.
- Backend mints Beep tokens via the Beep contract and updates MongoDB.
- Token transfers are executed on-chain, with off-chain records in MongoDB.
- Withdrawal converts Beep tokens to Naira, processed via the backend and bank integration.
- tAtom intent is recorded on-chain, enabling future exchange logic.

## How it Leverages Cosmos Technologies
Beep leverages Cosmos technologies through its deployment on Neutron and integration with the 
broader Cosmos ecosystem:

1. Neutron Blockchain (Cosmos-Based)
- Use Case: Hosts Beep’s smart contracts for tokenization, transfers, and intents.
- Implementation: Uses Neutron’s CosmWasm runtime for secure and scalable transaction logic.
- Benefit: Provides a neutral, interoperable platform optimized for financial applications.

2. CosmWasm
- Use Case: Powers the Beep token, Naira tokenization, and tAtom intent contracts.
- Implementation: Rust-based contracts deployed on Neutron, handling token minting, transfers, and 
exchange intents.
- Benefit: Enables programmable and secure financial operations on-chain.

3. Cosmos SDK (First iteration)
- Use Case: Facilitates blockchain integration from the backend.
- Implementation: Backend uses Cosmos SDK via cosmos.js to interact with Neutron contracts and 
process transactions.
- Benefit: Simplifies development and ensures compatibility with Cosmos standards.

4. ATOM (via tAtom)
- Use Case: Links Beep to the Cosmos Hub through tAtom exchange intents.
- Implementation: Token-Atom contract allows users to signal intent to swap Naira for tAtom, 
potentially bridging to ATOM.
- Benefit: Ties Beep to the Cosmos economy, enhancing interoperability.

## Future Plans and Roadmap
Beep aims to evolve into a robust mobile money solution with deeper Cosmos integration:

### Short-Term (Post-Hackathon, Q2 2025)

- Frontend Development: Build a fully functional mobile app for seamless user access.
- Bank Integration: Enhance Naira withdrawal with local bank APIs.
- Testing: Audit smart contracts and test on Neutron testnet for security and reliability.

### Medium-Term (Q3-Q4 2025)

- IBC Support: Enable cross-chain transfers of Beep tokens via Inter-Blockchain Communication (IBC).
- User Analytics: Add dashboards for transaction history and token balances.

### Long-Term (2026 and Beyond)

- Multi-Token Support: Expand to tokenize other fiat currencies or Cosmos tokens.
- Decentralized Governance: Introduce a governance model for Beep token holders.
- Ecosystem Growth: Partner with Cosmos projects and mobile money providers to scale adoption.

## Getting Started

To set up Beep locally, follow these steps:

### Prerequisites

- Node.js & npm: For running the backend.
- MongoDB: For off-chain data storage.
- Neutron Wallet: Fund with testnet NTRN (e.g., via Keplr).
- .env File: Configured with environment variables.

### Setup

Clone the Repository:

```bash
    git clone https://github.com/come-senusi-wale/beep-hackar.git
    cd beep-hacker
```

### Install Dependencies:

```bash
    npm install
```

### Configure Environment Variables:

Create a .env file:

```text
    PORT=5000
    MONGODB_URI=your_mongodb_url
    RPC=https://rpc-palvus.pion-1.ntrn.tech
    ENCRYPTION_KEY=9999
    TERMII_API_KEY=your_termii_api_key
    TERMII_SENDER_ID=your_termii_sender_id
    TOKEN_CONTRACT_ADDRESS=neutron1ujaf3dgpgn7e5tg6xy2hfnx6a6aupzjgxgj8fust08jttv03059s2jv2uw
    BEEP_CONTRACT_ADDRESS=neutron1rvn3dawaze4qyv7p9h4mcyll5axs0tefu90fdmw0s8u5wzaxjqhqc2m3gx
    TOKEN_ATOM_CONTRACT_ADDRESS=neutron17huqemgnu8r4092z74vu5jtzgm3lxg4gzqupu48648t8fz4wyzxsy7rjkf
    TOKEN_NGN_CONTRACT_ADDRESS=neutron1ujaf3dgpgn7e5tg6xy2hfnx6a6aupzjgxgj8fust08jttv03059s2jv2uw
    ADMIN_MNEMONIC=your_admin_mnemonic
```

### Run the Application:

```bash
    npm start
```

Backend runs at http://localhost:5000.

## Conclusion
Beep transforms mobile money by tokenizing Naira on the Neutron blockchain, leveraging Cosmos SDK 
and CosmWasm for a secure and scalable platform. With features like token transfers, Naira 
withdrawals, and tAtom intents, it bridges fiat and crypto economies. 

Explore the project at https://github.com/come-senusi-wale/beep-hackar
