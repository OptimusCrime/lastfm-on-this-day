import React from 'react';

import { getItem } from '../../utils/localStorage';
import { LocalStorageKeys } from '../../utils/localStorageKeys';
import { AuthenticationRequired, ListTracks } from './components';

export const Home = () => {
  const token = getItem(LocalStorageKeys.LOCAL_STORAGE_TOKEN_KEY);

  if (!token) {
    return <AuthenticationRequired />;
  }

  return <ListTracks accessToken={token} />;
};
