import ky from 'ky';

import { BackendEndpoints } from './backendEndpoints.types';

const api = ky.create({
  prefixUrl: process.env.NODE_ENV === 'production' ? 'https://lastfm.optimuscrime.net' : 'http://localhost:8113',
  retry: 0,
});

export const getAuth = () =>
  api
    .get('v1/auth')
    .json<BackendEndpoints.Auth.GET>()
    .then((res) => res);

export const postAuth = (token: string) =>
  api
    .post('v1/auth', {
      json: {
        token,
      }
    })
    .json<BackendEndpoints.Auth.POST>()
    .then((res) => res);

export const getTracks = (params: { token: string; date: string }) =>
  api
    .get(`v1/tracks?date=${params.date}`, {
      headers: {
        Authorization: `Bearer ${params.token}`
      }
    })
    .json<BackendEndpoints.Tracks.GET>()
    .then((res) => res);
