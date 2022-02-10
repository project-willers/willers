import { useState } from 'react'

/**
 * UseErrorTryFnCallback type.
 */
export type UseErrorTryFnCallback = () => Promise<void>

/**
 * UseErrorTryFn type.
 */
export type UseErrorTryFn = (fn: UseErrorTryFnCallback) => Promise<void>

/**
 * UseErrorReturn type.
 */
export type UseErrorReturn = [string | undefined, UseErrorTryFn]

/**
 * UseTransformer type.
 */
export type UseErrorTransformer = (e: unknown) => string | void

const defaultTransformer: UseErrorTransformer = (e) => {
  if (e instanceof Error) {
    return e.message
  }
}

/**
 * use error.
 */
export const useError = (
  transform: UseErrorTransformer = defaultTransformer
): UseErrorReturn => {
  const [error, setError] = useState<string>()

  const tryFn: UseErrorTryFn = async (callback) => {
    setError(undefined)

    try {
      await callback()
    } catch (e) {
      setError(transform(e) ?? 'エラーが発生しました')
    }
  }

  return [error, tryFn]
}
