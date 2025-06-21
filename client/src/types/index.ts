// 共通の型定義

export interface User {
  id: string
  name: string
  email?: string
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
  rating?: number
  types?: string[]
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

// 場所検索の設定
export interface PlaceSearchSettings {
  genre: string
  baseLocation: string
  useCurrentLocation: boolean
  radius: number
  maxCandidates: number
}

// フォームデータ
export interface DateTimeFormData {
  date: string
  startTime: string
  endTime: string
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
