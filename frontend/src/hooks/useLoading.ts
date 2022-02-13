import { useCallback, useState } from 'react'

/**
 * UseLoaidngLoaderCallback type.
 */
export type UseLoadingLoaderCallback = () => Promise<void>

/**
 * UseLoadingLoaderFn type.
 */
export type UseLoadingLoaderFn = (
  callback: UseLoadingLoaderCallback
) => Promise<void>

/**
 * UseLoadingReturn type.
 */
export type UseLoadingReturn = [boolean, UseLoadingLoaderFn]

/**
 * use loading.
 */
export const useLoading = (): UseLoadingReturn => {
  const [loading, setLoading] = useState(false)

  const loader = useCallback(async (callback: UseLoadingLoaderCallback) => {
    setLoading(true)

    await callback()

    setLoading(false)
  }, [])

  return [loading, loader]
}
