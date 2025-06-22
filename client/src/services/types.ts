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

// Room関連のAPI型定義
export interface TimeOption {
  start_time: string
  end_time: string
}

export interface ReturnTimeOption {
  id: string
  start_time: string
  end_time: string
}

export interface ReturnPlaceOption {
  id: string
  google_place_id: string
  name: string
  location: {
    lat: number
    lng: number
  }
  photoReference: string
  priceLevel: number
  rating: number
  address: string
}

export interface NewCreateRoomRequest {
  name: string
  time_options: TimeOption[]
  place_options: string[]
  center_place_id: string
  radius: number
  place_max: number
}

export interface NewCreateRoomResponse {
  room_id: string
  time_options: ReturnTimeOption[]
  place_options: ReturnPlaceOption[]
}

export interface GetRoomResponse {
  room_id: string
  name: string
  center_point: string
  radius: number
  place_max: number
  time_options: ReturnTimeOption[]
  place_options: ReturnPlaceOption[]
}

// Vote関連のAPI型定義
export interface PlaceVoteData {
  id: string
  rank: number
}

export interface PostVoteRequest {
  roomID: string
  name: string
  timeId: string[]
  placeData: PlaceVoteData[]
}

export interface PostVoteResponse {
  userID: string
}

// 投票結果関連の型定義
export interface UserNameID {
  name: string
  userID: string
}

export interface TimeResult {
  startTime: string
  endTime: string
  user: UserNameID[]
}

export interface PlaceResult {
  googlePlaceID: string
  placeName: string
  point: number
}

export interface VoteResultResponse {
  place: PlaceResult[]
  time: TimeResult[]
}
