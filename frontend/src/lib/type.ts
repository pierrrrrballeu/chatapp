enum MessageType {
  Text,
  Join,
  Leave,
}

type ChatMessage =  {
  content: string
  timestamp: string
  username: string
  type: MessageType
}

export { type ChatMessage, MessageType }