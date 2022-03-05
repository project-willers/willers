import { User } from '@/types/User'
import { atom } from 'jotai'
import decode from 'jwt-decode'

export const jwtAtom = atom<string | null>(null)

export const userAtom = atom((get) => {
  const jwt = get(jwtAtom)

  if (jwt) {
    return decode<User>(jwt)
  }

  return null
})
