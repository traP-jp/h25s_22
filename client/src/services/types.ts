// API関連の型定義
export interface PlaceSearchRequest {
  name: string
  radius: number
  type: string
  maxSize: number
  language: string
}

export interface PlaceSearchResult {
  name: string
  placeID: string
  location: {
    lat: number
    lng: number
  }
  photoReference: string
  priceLevel?: number
  rating?: number
  address?: string
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
