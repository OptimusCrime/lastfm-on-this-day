import { LocalStorageKeys } from './localStorageKeys';

export const getItem = (key: LocalStorageKeys): string | null => {
  if (!localStorage) {
    return null;
  }

  return localStorage.getItem(key) ?? null;
};

export const setItem = (key: LocalStorageKeys, value: string): void => {
  if (!localStorage) {
    return;
  }

  localStorage.setItem(key, value);
};
