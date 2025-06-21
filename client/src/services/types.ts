// API関連の型定義
export interface PlaceSearchRequest {
  name: string
  radius: number
  type: string
  maxSize: number
  language: string
}

export interface PlaceSearchResult {
  placeID: string
  name: string
  address: string
  location: {
    lat: number
    lng: number
  }
  photoReference?: {
    photo_reference: string
    height: number
    width: number
    html_attributions: string[]
  }
  priceLevel?: number
  rating?: number
  types?: string[]
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
