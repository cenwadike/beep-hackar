import { Event, Attribute } from '@cosmjs/stargate'
import { Tendermint34Client, WebsocketClient } from '@cosmjs/tendermint-rpc'

// Event type constants
const EVENT_TYPES = {
  CREATE_INTENT: 'create_intent',
  ESCROW_IBC_TRANSFER: 'escrow_ibc_transfer',
  ESCROW_NATIVE_TOKENS: 'escrow_native_tokens',
  IBC_TRANSFER_FAILURE: 'ibc_transfer_failure',
  ACCEPT_INTENT: 'accept_intent',
  ACCEPT_INTENT_FAILURE: 'accept_intent_failure',
  EXECUTOR_ESCROW: 'executor_escrow',
  CREATOR_PAYOUT: 'creator_payout',
  EXECUTOR_PAYOUT: 'executor_payout',
  INTENT_EXECUTED: 'intent_executed'
} as const

// Interface definitions for each event type
interface CreateIntentEvent {
  intentId: string
  creator: string
  amount: string
  inputToken: string
  outputToken: string
  targetChain: string
  minOutput: string
  status: string
  expiryHeight: string
}

interface EscrowIBCTransferEvent {
  intentId: string
  sender: string
  receiver: string
  ibcDenom: string
  amount: string
}

interface EscrowNativeTokensEvent {
  intentId: string
  sender: string
  receiver: string
  denom: string
  amount: string
}

interface IBCTransferFailureEvent {
  sender: string
  receiver: string
  denom: string
  amount: string
  error: string
}

interface AcceptIntentEvent {
  intentId: string
  creator: string
  executor: string
  inputToken: string
  outputToken: string
  amount: string
  minOutput: string
}

interface AcceptIntentFailureEvent {
  intentId: string
  error: string
  stage: string
}

interface ExecutorEscrowEvent {
  intentId: string
  sender: string
  receiver: string
  denom?: string
  ibcDenom?: string
  amount: string
}

interface CreatorPayoutEvent {
  intentId: string
  sender: string
  receiver: string
  denom?: string
  ibcDenom?: string
  amount: string
}

interface ExecutorPayoutEvent {
  intentId: string
  sender: string
  receiver: string
  denom?: string
  ibcDenom?: string
  amount: string
}

interface IntentExecutedEvent {
  intentId: string
  creator: string
  executor: string
  status: string
}

// Helper function to parse attributes
function parseAttributes(attributes: readonly Attribute[]): Record<string, string> {
  return attributes.reduce((acc, attr) => ({
    ...acc,
    [attr.key]: attr.value
  }), {})
}

// Main event handler class
export class IntentEventHandler {
  private onCreateIntent?: (event: CreateIntentEvent) => void
  private onEscrowIBCTransfer?: (event: EscrowIBCTransferEvent) => void
  private onEscrowNativeTokens?: (event: EscrowNativeTokensEvent) => void
  private onIBCTransferFailure?: (event: IBCTransferFailureEvent) => void
  private onAcceptIntent?: (event: AcceptIntentEvent) => void
  private onAcceptIntentFailure?: (event: AcceptIntentFailureEvent) => void
  private onExecutorEscrow?: (event: ExecutorEscrowEvent) => void
  private onCreatorPayout?: (event: CreatorPayoutEvent) => void
  private onExecutorPayout?: (event: ExecutorPayoutEvent) => void
  private onIntentExecuted?: (event: IntentExecutedEvent) => void

  handleEvent(event: Event): void {
    const attributes = parseAttributes(event.attributes)

    switch (event.type) {
      case EVENT_TYPES.CREATE_INTENT:
        if (this.onCreateIntent) {
          this.onCreateIntent({
            intentId: attributes.intent_id,
            creator: attributes.creator,
            amount: attributes.amount,
            inputToken: attributes.input_token,
            outputToken: attributes.output_token,
            targetChain: attributes.target_chain,
            minOutput: attributes.min_output,
            status: attributes.status,
            expiryHeight: attributes.expiry_height
          })
        }
        break

      case EVENT_TYPES.ACCEPT_INTENT:
        if (this.onAcceptIntent) {
          this.onAcceptIntent({
            intentId: attributes.intent_id,
            creator: attributes.creator,
            executor: attributes.executor,
            inputToken: attributes.input_token,
            outputToken: attributes.output_token,
            amount: attributes.amount,
            minOutput: attributes.min_output
          })
        }
        break

      case EVENT_TYPES.ACCEPT_INTENT_FAILURE:
        if (this.onAcceptIntentFailure) {
          this.onAcceptIntentFailure({
            intentId: attributes.intent_id,
            error: attributes.error,
            stage: attributes.stage
          })
        }
        break

      case EVENT_TYPES.EXECUTOR_ESCROW:
        if (this.onExecutorEscrow) {
          this.onExecutorEscrow({
            intentId: attributes.intent_id,
            sender: attributes.sender,
            receiver: attributes.receiver,
            denom: attributes.denom,
            ibcDenom: attributes.ibc_denom,
            amount: attributes.amount
          })
        }
        break

      case EVENT_TYPES.CREATOR_PAYOUT:
        if (this.onCreatorPayout) {
          this.onCreatorPayout({
            intentId: attributes.intent_id,
            sender: attributes.sender,
            receiver: attributes.receiver,
            denom: attributes.denom,
            ibcDenom: attributes.ibc_denom,
            amount: attributes.amount
          })
        }
        break

      case EVENT_TYPES.EXECUTOR_PAYOUT:
        if (this.onExecutorPayout) {
          this.onExecutorPayout({
            intentId: attributes.intent_id,
            sender: attributes.sender,
            receiver: attributes.receiver,
            denom: attributes.denom,
            ibcDenom: attributes.ibc_denom,
            amount: attributes.amount
          })
        }
        break

      case EVENT_TYPES.INTENT_EXECUTED:
        if (this.onIntentExecuted) {
          this.onIntentExecuted({
            intentId: attributes.intent_id,
            creator: attributes.creator,
            executor: attributes.executor,
            status: attributes.status
          })
        }
        break

      // Existing handlers from previous implementation
      case EVENT_TYPES.ESCROW_IBC_TRANSFER:
        if (this.onEscrowIBCTransfer) {
          this.onEscrowIBCTransfer({
            intentId: attributes.intent_id,
            sender: attributes.sender,
            receiver: attributes.receiver,
            ibcDenom: attributes.ibc_denom,
            amount: attributes.amount
          })
        }
        break

      case EVENT_TYPES.ESCROW_NATIVE_TOKENS:
        if (this.onEscrowNativeTokens) {
          this.onEscrowNativeTokens({
            intentId: attributes.intent_id,
            sender: attributes.sender,
            receiver: attributes.receiver,
            denom: attributes.denom,
            amount: attributes.amount
          })
        }
        break

      case EVENT_TYPES.IBC_TRANSFER_FAILURE:
        if (this.onIBCTransferFailure) {
          this.onIBCTransferFailure({
            sender: attributes.sender,
            receiver: attributes.receiver,
            denom: attributes.denom,
            amount: attributes.amount,
            error: attributes.error
          })
        }
        break
    }
  }

  // Event registration methods
  setCreateIntentHandler(handler: (event: CreateIntentEvent) => void): void {
    this.onCreateIntent = handler
  }

  setAcceptIntentHandler(handler: (event: AcceptIntentEvent) => void): void {
    this.onAcceptIntent = handler
  }

  setAcceptIntentFailureHandler(handler: (event: AcceptIntentFailureEvent) => void): void {
    this.onAcceptIntentFailure = handler
  }

  setExecutorEscrowHandler(handler: (event: ExecutorEscrowEvent) => void): void {
    this.onExecutorEscrow = handler
  }

  setCreatorPayoutHandler(handler: (event: CreatorPayoutEvent) => void): void {
    this.onCreatorPayout = handler
  }

  setExecutorPayoutHandler(handler: (event: ExecutorPayoutEvent) => void): void {
    this.onExecutorPayout = handler
  }

  setIntentExecutedHandler(handler: (event: IntentExecutedEvent) => void): void {
    this.onIntentExecuted = handler
  }

  setEscrowIBCTransferHandler(handler: (event: EscrowIBCTransferEvent) => void): void {
    this.onEscrowIBCTransfer = handler
  }

  setEscrowNativeTokensHandler(handler: (event: EscrowNativeTokensEvent) => void): void {
    this.onEscrowNativeTokens = handler
  }

  setIBCTransferFailureHandler(handler: (event: IBCTransferFailureEvent) => void): void {
    this.onIBCTransferFailure = handler
  }
}


export class IntentEventListener {
    private client: Tendermint34Client | null = null
    private eventHandler: IntentEventHandler
    private wsUrl: string
    private subscriptions: string[] = []
    private reconnectAttempts = 0
    private readonly maxReconnectAttempts = 5
    private readonly reconnectDelay = 1000 // 1 second
  
    constructor(wsUrl: string) {
      this.wsUrl = wsUrl
      this.eventHandler = new IntentEventHandler()
    }
  
    async connect(): Promise<void> {
      try {
        const wsClient = new WebsocketClient(this.wsUrl)
        this.client = await Tendermint34Client.create(wsClient)
        this.reconnectAttempts = 0
        console.log('Connected to blockchain at', this.wsUrl)
        
        // Subscribe to all intent-related events
        await this.subscribeToEvents()
      } catch (error) {
        console.error('Connection error:', error)
        await this.handleReconnect()
      }
    }
  
    private async handleReconnect(): Promise<void> {
      if (this.reconnectAttempts < this.maxReconnectAttempts) {
        this.reconnectAttempts++
        console.log(`Reconnection attempt ${this.reconnectAttempts}/${this.maxReconnectAttempts}`)
        await new Promise(resolve => setTimeout(resolve, this.reconnectDelay))
        await this.connect()
      } else {
        throw new Error('Failed to connect after maximum reconnection attempts')
      }
    }
  
    private async subscribeToEvents(): Promise<void> {
      if (!this.client) {
        throw new Error('Client not initialized')
      }
  
      // Create query strings for each event type
      const eventQueries = Object.values(EVENT_TYPES).map(eventType => `event.type='${eventType}'`)
  
      try {
        // Subscribe to events
        for (const query of eventQueries) {
          const stream = this.client.subscribeNewBlock()
          
          // Process events from each block
          stream.subscribe({
            next: async (event) => {
              if (event.header) {
                // Fetch the block data
                const block = await this.client!.block(event.header.height)
                
                // Process all transactions in the block
                if (block.block.txs) {
                  for (const tx of block.block.txs) {
                    // Get transaction results
                    const txResult = await this.client!.tx({ hash: tx })
                    
                    // Process events from transaction
                    if (txResult.result.events) {
                      txResult.result.events
                        .filter((event: { type: any }) => Object.values(EVENT_TYPES).includes(event.type as any))
                        .forEach((event) => {
                          this.eventHandler.handleEvent({
                            type: event.type,
                            attributes: event.attributes.map(attr => ({
                              key: Buffer.from(attr.key).toString('utf8'),
                              value: Buffer.from(attr.value).toString('utf8')
                            }))
                          })
                        })
                    }
                  }
                }
              }
            },
            error: async (error: any) => {
              console.error('Subscription error:', error)
              await this.handleReconnect()
            },
            complete: () => {
              console.log('Subscription completed')
            }
          })
          
          this.subscriptions.push(query)
        }
    
        console.log('Subscribed to events:', this.subscriptions)
    } catch (error) {
      console.error('Subscription error:', error)
      throw error
    }
  }

  async disconnect(): Promise<void> {
    if (this.client) {
        this.client.disconnect()
        this.client = null
        this.subscriptions = []
        console.log('Disconnected from blockchain')
    }
  }

  // Event handler setters
  setCreateIntentHandler(handler: (event: CreateIntentEvent) => void): void {
    this.eventHandler.setCreateIntentHandler(handler)
  }

  setAcceptIntentHandler(handler: (event: AcceptIntentEvent) => void): void {
    this.eventHandler.setAcceptIntentHandler(handler)
  }
  setAcceptIntentFailureHandler(handler: (event: AcceptIntentFailureEvent) => void): void {
    this.eventHandler.setAcceptIntentFailureHandler(handler)
  }

  setExecutorEscrowHandler(handler: (event: ExecutorEscrowEvent) => void): void {
    this.eventHandler.setExecutorEscrowHandler(handler)
  }

  setCreatorPayoutHandler(handler: (event: CreatorPayoutEvent) => void): void {
    this.eventHandler.setCreatorPayoutHandler(handler)
  }

  setExecutorPayoutHandler(handler: (event: ExecutorPayoutEvent) => void): void {
    this.eventHandler.setExecutorPayoutHandler(handler)
  }

  setIntentExecutedHandler(handler: (event: IntentExecutedEvent) => void): void {
    this.eventHandler.setIntentExecutedHandler(handler)
  }

  setEscrowIBCTransferHandler(handler: (event: EscrowIBCTransferEvent) => void): void {
    this.eventHandler.setEscrowIBCTransferHandler(handler)
  }

  setEscrowNativeTokensHandler(handler: (event: EscrowNativeTokensEvent) => void): void {
    this.eventHandler.setEscrowNativeTokensHandler(handler)
  }

  setIBCTransferFailureHandler(handler: (event: IBCTransferFailureEvent) => void): void {
    this.eventHandler.setIBCTransferFailureHandler(handler)
  }
}

// Example usage
const startEventListener = async () => {
  const listener = new IntentEventListener('ws://your-cosmos-node:26657/websocket')

  // Set up event handlers
  listener.setCreateIntentHandler((event) => {
    console.log('New intent created:', event)
  })

  listener.setAcceptIntentHandler((event) => {
    console.log('Intent accepted:', event)
  })

  // Connect to the blockchain
  try {
    await listener.connect()
  } catch (error) {
    console.error('Failed to start listener:', error)
  }

  // Cleanup on process termination
  process.on('SIGINT', async () => {
    console.log('Shutting down listener...')
    await listener.disconnect()
    process.exit(0)
  })
}

// Start the listener
startEventListener().catch(console.error)  