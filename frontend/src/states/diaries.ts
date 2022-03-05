import { Diary } from '@/types/Diary'
import { atom } from 'jotai'

export const diariesAtom = atom<Diary[]>([])

export const editDiaryAtom = atom<Diary | null>(null)
