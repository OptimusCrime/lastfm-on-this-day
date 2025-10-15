import React, { useEffect, useState } from 'react';

import { getAuth, postAuth } from '../../../../api/endpoints/backendEndpoints';
import { Card } from '../../../../layout/Card';
import { setItem } from '../../../../utils/localStorage';
import { LocalStorageKeys } from '../../../../utils/localStorageKeys';

interface AuthenticationRequiredState {
  isLoading: boolean;
  error: boolean;
  url: string | null;
}

export const AuthenticationRequired = () => {
  const [state, setState] = useState<AuthenticationRequiredState>({
    isLoading: true,
    error: false,
    url: null,
  });

  const getAuthCall = async () => {
    try {
      const data = await getAuth();

      setState({
        isLoading: false,
        error: false,
        url: data.url,
      });
    } catch (_) {
      setState({
        isLoading: false,
        error: true,
        url: null,
      });
    }
  };

  const postAuthCall = async (token: string) => {
    try {
      const data = await postAuth(token);

      setItem(LocalStorageKeys.LOCAL_STORAGE_TOKEN_KEY, data.accessToken);

      window.location.href = window.location.href.split('?')[0];
    } catch (_) {
      setState({
        isLoading: false,
        error: true,
        url: null,
      });
    }
  };

  useEffect(() => {
    const searchParams = new URLSearchParams(window.location.search);
    const token = searchParams.get('token');

    if (token) {
      postAuthCall(token);
    } else {
      getAuthCall();
    }
  }, []);

  if (state.url) {
    return (
      <Card>
        <div className="flex justify-center">
          <a href={state.url} className="btn">
            Sign in to Last.fm
          </a>
        </div>
      </Card>
    );
  }

  if (state.error) {
    return (
      <Card>
        <div className="flex justify-center">
          <p>Woops. Something went wrong!</p>
        </div>
      </Card>
    );
  }

  return (
    <Card>
      <div className="flex justify-center">
        <p>Please wait&hellip;</p>
      </div>
    </Card>
  );
};
