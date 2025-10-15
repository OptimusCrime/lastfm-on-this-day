import { Track } from '../../types/tracks';

/* eslint-disable @typescript-eslint/no-namespace */
export namespace BackendEndpoints {
  export namespace Auth {
    export interface GET {
      url: string;
    }

    export interface POST {
      accessToken: string;
    }
  }

  export namespace Tracks {
    export interface GET {
      data: Track[];
    }
  }
}
