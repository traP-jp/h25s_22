// 共通の型定義

export interface User {
  id: string
  name: string
  email?: string
}

export interface Room {
  id: string
  name: string
  description: string
  createdAt: string
  timeOptions: TimeOption[]
  places: Place[]
  participants?: User[]
}

export interface TimeOption {
  id: string
  date: string
  startTime: string
  endTime: string
}

export interface Place {
  id: string
  name: string
  address: string
  latitude?: number
  longitude?: number
  description?: string
  imageUrl?: string
}

export interface PlaceVote {
  id: string
  placeId: string
  userId: string
  score: number
  comment?: string
}

export interface TimeVote {
  id: string
  timeOptionId: string
  userId: string
  available: boolean
}

// API レスポンス型
export interface ApiResponse<T> {
  data: T
  message?: string
  success: boolean
}

// エラー型
export interface ApiError {
  message: string
  code?: string
  details?: Record<string, any>
}
