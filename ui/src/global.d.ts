interface NetworkError {
  statusCode: number
  statusText: string
  response: Response
  message?: string
}
