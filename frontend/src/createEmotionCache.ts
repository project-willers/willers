import createCache from '@emotion/cache'

/**
 * create emotion cache.
 */
export const createEmotionCache = () => {
  return createCache({ key: 'css', prepend: true })
}
