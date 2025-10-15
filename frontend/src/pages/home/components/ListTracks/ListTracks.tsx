import React, {useState} from "react";
import {Track} from "../../../../types/tracks";
import {getTracks} from "../../../../api/endpoints/backendEndpoints";
import {addLeadingZero} from "../../../../utils/strings";

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

  for (let year = new Date().getFullYear(); year >= 2005 ; year--) {
    state.push({
      year,
      hasFetched: false,
      isLoading: false,
      error: false,
      tracks: []
    })
  }

  return state;
}

const constructDate = ({selectedDate, year}: {selectedDate: Date, year: number}): string =>
  `${year}-${addLeadingZero(selectedDate.getMonth() + 1)}-${addLeadingZero(selectedDate.getDate())}`;


export const ListTracks = (props: Props) => {
  const [date, setDate] = useState<Date>(new Date());
  const [state, setState] = React.useState<State[]>(initState());

  const getTracksCall = async (year: number) => {
    setState(prevState => {
      return prevState.map(item => {
        if (item.year !== year) {
          return item;
        }

        return {
          ...item,
          isLoading: true,
          error: false
        }
      })
    });

    try {
      const data = await getTracks({
        token: props.accessToken,
        date: constructDate({
          selectedDate: date,
          year
        })
      });

      setState(prevState => {
        return prevState.map(item => {
          if (item.year !== year) {
            return item;
          }

          return {
            ...item,
            hasFetched: true,
            tracks: data,
            isLoading: false,
            error:false
          }
        })
      });
    }
    catch(_) {
      setState(prevState => {
        return prevState.map(item => {
          if (item.year !== year) {
            return item;
          }

          return {
            ...item,
            hasFetched: true,
            isLoading: false,
            error: true
          }
        })
      });
    }
  }

  return (
    <div>
      {date.getDate()}
      {state.map(item => (
        <div key={item.year}>
          <h2>{item.year}</h2>

          {!item.hasFetched && (
            <button
              className='btn'
              onClick={async () => await getTracksCall(item.year)}
              disabled={item.isLoading}
            >
              {item.isLoading ? 'Loading...' : 'Load tracks'}
            </button>
          )}
        </div>
      ))}
    </div>
  );
}
