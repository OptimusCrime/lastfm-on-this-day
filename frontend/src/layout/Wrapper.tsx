import React from 'react';

import { Home } from '../pages';
import { Header } from './Header';

export const Wrapper = () => (
  <div className="container max-w-none mb-16">
    <div className="container max-w-none bg-neutral">
      <div className="container mb-8">
        <Header />
      </div>
    </div>
    <div className="container mx-auto">
      <Home />
    </div>
  </div>
);
