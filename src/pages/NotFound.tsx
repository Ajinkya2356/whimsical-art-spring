
import React from 'react';
import { Button } from '@/components/ui/button';
import { useLocation } from 'react-router-dom';
import { useEffect } from 'react';

const NotFound = () => {
  const location = useLocation();

  useEffect(() => {
    console.error(
      "404 Error: User attempted to access non-existent route:",
      location.pathname
    );
  }, [location.pathname]);

  return (
    <div className="min-h-screen flex flex-col items-center justify-center text-center p-6 bg-ghibli-gradient">
      <div className="relative mb-8 animate-float">
        <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-24 h-24 bg-white rounded-full opacity-80"></div>
        <h1 className="text-8xl font-pacifico text-ghibli-forest relative z-10">404</h1>
      </div>
      
      <h2 className="text-2xl font-pacifico text-ghibli-forest mb-4">Spirited Away</h2>
      <p className="text-lg text-ghibli-forest/80 max-w-md mb-8">
        It seems you've wandered into a realm that doesn't exist. Let's find our way back.
      </p>
      
      <Button className="ghibli-button" asChild>
        <a href="/">Return to the Garden</a>
      </Button>
    </div>
  );
};

export default NotFound;
