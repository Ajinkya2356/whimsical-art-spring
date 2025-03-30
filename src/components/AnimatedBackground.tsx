
import React from 'react';

const AnimatedBackground = () => {
  return (
    <div className="fixed inset-0 overflow-hidden pointer-events-none z-0">
      <div className="cloud w-40 h-20 left-[10%] top-[15%]"></div>
      <div className="cloud w-56 h-24 right-[20%] top-[10%] opacity-80 animate-float" style={{ animationDelay: '1s' }}></div>
      <div className="cloud w-32 h-16 left-[30%] top-[25%] opacity-70 animate-float" style={{ animationDelay: '2s' }}></div>
      <div className="cloud w-48 h-20 right-[15%] top-[30%] opacity-60 animate-float" style={{ animationDelay: '3s' }}></div>
      <div className="cloud w-36 h-16 left-[15%] top-[40%] opacity-60 animate-float" style={{ animationDelay: '4s' }}></div>
      
      <div className="absolute bottom-0 left-0 right-0 h-32 bg-gradient-to-t from-ghibli-grass/20 to-transparent"></div>
    </div>
  );
};

export default AnimatedBackground;
