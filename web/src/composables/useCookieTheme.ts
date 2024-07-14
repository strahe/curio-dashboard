import { useCookies } from '@vueuse/integrations/useCookies'

export function useCookieTheme() {
  const theme = useCookies(['theme'])
  if (theme.get('theme') === undefined){
    theme.set('theme', 'system');
  }
  return theme;
}
