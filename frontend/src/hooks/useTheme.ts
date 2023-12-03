import { useMantineColorScheme } from '@mantine/core';

function useTheme() {
  // -> colorScheme is 'auto' | 'light' | 'dark'
  return useMantineColorScheme({keepTransitions: true});
}

export default useTheme;