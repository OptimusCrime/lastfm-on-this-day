import React from 'react';

import {getItem} from "../../utils/localStorage";
import {LocalStorageKeys} from "../../utils/localStorageKeys";
import {AuthenticationRequired, ListTracks} from "./components";

export const Home = () => {
  const token = getItem(LocalStorageKeys.LOCAL_STORAGE_TOKEN_KEY)

  return (
    <div className="flex justify-center">
      <div className="flex w-full max-w-[900px] px-8">
        <div className="card bg-neutral text-neutral-content card-compact w-full">
          <div className="card-body flex justify-start">
            {token === null ? <AuthenticationRequired /> : <ListTracks accessToken={token} />}
          </div>
        </div>
      </div>
    </div>
  );
}
