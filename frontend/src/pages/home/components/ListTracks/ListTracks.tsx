import React, { useState } from 'react';

import { getTracks } from '../../../../api/endpoints/backendEndpoints';
import { CaretLeftIcon } from '../../../../icons/CaretLeftIcon';
import { CaretRightIcon } from '../../../../icons/CaretRightIcon';
import { Card } from '../../../../layout/Card';
import { Track } from '../../../../types/tracks';
import { addLeadingZero } from '../../../../utils/strings';

interface Props {
  accessToken: string;
}

interface State {
  year: number;
  hasFetched: boolean;
  isLoading: boolean;
  error: boolean;
  tracks: Track[];
}

const initState = () => {
  const state: State[] = [];

  for (let year = new Date().getFullYear(); year >= 2005; year--) {
    state.push({
      year,
      hasFetched: false,
      isLoading: false,
      error: false,
      tracks: [],
    });
  }

  return state;
};

const getMonthName = (month: number) => {
  switch (month) {
    case 0:
      return 'January';
    case 1:
      return 'February';
    case 2:
      return 'March';
    case 3:
      return 'April';
    case 4:
      return 'May';
    case 5:
      return 'June';
    case 6:
      return 'July';
    case 7:
      return 'August';
    case 8:
      return 'September';
    case 9:
      return 'October';
    case 10:
      return 'November';
    default:
      return 'December';
  }
};

const constructDate = ({ selectedDate, year }: { selectedDate: Date; year: number }): string =>
  `${year}-${addLeadingZero(selectedDate.getMonth() + 1)}-${addLeadingZero(selectedDate.getDate())}`;

const sumTotalPlays = (tracks: Track[]) => {
  let sum = 0;

  for (const track of tracks) {
    sum += track.playCount;
  }

  return sum;
};

export const ListTracks = (props: Props) => {
  const [date, setDate] = useState<Date>(new Date());
  const [state, setState] = React.useState<State[]>(initState());

  const getTracksCall = async (year: number) => {
    setState((prevState) => {
      return prevState.map((item) => {
        if (item.year !== year) {
          return item;
        }

        return {
          ...item,
          isLoading: true,
          error: false,
        };
      });
    });

    try {
      const response = await getTracks({
        token: props.accessToken,
        date: constructDate({
          selectedDate: date,
          year,
        }),
      });

      setState((prevState) => {
        return prevState.map((item) => {
          if (item.year !== year) {
            return item;
          }

          return {
            ...item,
            hasFetched: true,
            tracks: response.data,
            isLoading: false,
            error: false,
          };
        });
      });
    } catch (_) {
      setState((prevState) => {
        return prevState.map((item) => {
          if (item.year !== year) {
            return item;
          }

          return {
            ...item,
            hasFetched: true,
            isLoading: false,
            error: true,
          };
        });
      });
    }
  };

  return (
    <div>
      <div className="flex justify-center mb-4">
        <div className="flex w-full max-w-[650px] px-4">
          <div className="card w-full">
            <div className="card-body flex justify-center">
              <div className="flex flex-row justify-center">
                <div className="flex bg-neutral text-center justify-center items-center border-neutral flex-row p-4 w-full max-w-[400px] rounded-lg">
                  <div
                    className="w-[48px] h-[48px] flex justify-center items-center rounded-full cursor-pointer fill-neutral-content"
                    title="Previous day"
                    onClick={() => {
                      setDate((prevState) => {
                        const newDate = new Date(prevState);
                        newDate.setDate(newDate.getDate() - 1);
                        return newDate;
                      });

                      setState(initState());
                    }}
                  >
                    <CaretLeftIcon />
                  </div>
                  <div className="flex flex-col px-4 max-w-[400px] w-full">
                    <div>
                      <span className="text-lg">{date.getDate()}.</span>
                    </div>
                    <div>
                      <span className="text-md">
                        <strong>{getMonthName(date.getMonth())}</strong>
                      </span>
                    </div>
                  </div>
                  <div
                    className="w-[48px] h-[48px] flex justify-center items-center rounded-full cursor-pointer fill-neutral-content"
                    title="Next day"
                    onClick={() => {
                      setDate((prevState) => {
                        const newDate = new Date(prevState);
                        newDate.setDate(newDate.getDate() + 1);

                        return newDate;
                      });

                      setState(initState());
                    }}
                  >
                    <CaretRightIcon />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      {state.map((item) => (
        <Card key={item.year}>
          <div>
            <h4 className="text-xl pb-2">{item.year}</h4>

            {item.hasFetched &&
              (item.tracks.length === 0 ? (
                <p className="pb-2">No tracks played on this date this year. Try another year.</p>
              ) : (
                <p className="pb-2">
                  Total plays: <strong>{sumTotalPlays(item.tracks)}</strong>
                </p>
              ))}

            {item.tracks && (
              <ul className="list-disc list-inside">
                {item.tracks.map((track, idx) => (
                  <li key={idx}>
                    [{track.playCount}] <strong>{track.artist ?? 'Unknown artist'}</strong>:{' '}
                    {track.name ?? 'Unknown track'}
                    {track.album ? ` (${track.album})` : ''}
                  </li>
                ))}
              </ul>
            )}

            {!item.hasFetched && (
              <button className="btn" onClick={async () => await getTracksCall(item.year)} disabled={item.isLoading}>
                {item.isLoading ? 'Loading...' : 'Load tracks'}
              </button>
            )}
          </div>
        </Card>
      ))}
    </div>
  );
};
