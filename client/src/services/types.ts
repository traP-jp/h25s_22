// API関連の型定義
export interface PlaceSearchRequest {
  name: string
  radius: number
  type: string
  maxSize: number
  language: string
}

export interface PlaceSearchResult {
  id: string
  name: string
  address: string
  latitude: number
  longitude: number
  rating?: number
  types: string[]
  photoUrl?: string
}

export interface CreateRoomRequest {
  name: string
  description: string
  timeOptions: {
    date: string
    startTime: string
    endTime: string
  }[]
  places: PlaceSearchResult[]
}

export interface CreateRoomResponse {
  id: string
  message: string
  success: boolean
}
