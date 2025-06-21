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
    const response = await fetch(`${API_BASE_URL}/place/nearSearch`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(params),
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const data = await response.json()
    return data.results || []
  } catch (error) {
    console.error('Place search API error:', error)
    throw error
  }
}

// ルーム作成API（仮実装）
export const createRoom = async (roomData: CreateRoomRequest): Promise<CreateRoomResponse> => {
  try {
    // 実際のAPIが完成するまでの仮実装
    console.log('Creating room with data:', roomData)

    // 仮のレスポンスを返す
    await new Promise((resolve) => setTimeout(resolve, 1000)) // API呼び出しをシミュレート

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
