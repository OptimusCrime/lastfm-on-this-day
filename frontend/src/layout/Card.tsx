import React from 'react';

interface Props {
  children: React.ReactNode;
}

export const Card = (props: Props) => (
  <div className="flex justify-center mb-4">
    <div className="flex w-full max-w-[650px] px-4">
      <div className="card bg-neutral text-neutral-content card-compact w-full">
        <div className="card-body flex justify-start">{props.children}</div>
      </div>
    </div>
  </div>
);
