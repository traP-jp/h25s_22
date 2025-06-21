import type {
  PlaceSearchRequest,
  PlaceSearchResult,
  CreateRoomRequest,
  CreateRoomResponse,
} from './types'

const API_BASE_URL = 'http://localhost:8080/api/v1'

// 場所検索API
export const searchNearbyPlaces = async (
  params: PlaceSearchRequest,
): Promise<PlaceSearchResult[]> => {
  try {
    const queryParams = new URLSearchParams({
      name: params.name,
      radius: params.radius.toString(),
      type: params.type,
      maxSize: params.maxSize.toString(),
      language: params.language,
    })

    const response = await fetch(`${API_BASE_URL}/place/nearSearch?${queryParams}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const data = await response.json()
    return data || []
  } catch (error) {
    console.error('Place search API error:', error)
    throw error
  }
}

// ルーム作成API（仮実装）
export const createRoom = async (roomData: CreateRoomRequest): Promise<CreateRoomResponse> => {
  try {
    // TODO: 実際のAPIエンドポイントに置き換える
    console.log('Creating room with data:', roomData)

    // 仮のレスポンスを返す
    await new Promise((resolve) => setTimeout(resolve, 1000))

    return {
      id: `room_${Date.now()}`,
      message: 'ルームが正常に作成されました',
      success: true,
    }
  } catch (error) {
    console.error('Room creation API error:', error)
    throw error
  }
}
