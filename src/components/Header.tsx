
import React from 'react';

const Header = () => {
  return (
    <header className="w-full py-4 px-6 md:px-10 flex justify-between items-center gap-4 relative z-10">
      <div className="flex items-center">
        <h1 className="text-3xl md:text-4xl bg-gradient-to-r from-ghibli-accent to-ghibli-sunset bg-clip-text text-transparent py-3">Prompt Garden</h1>
      </div>
    </header>
  );
};

export default Header;
