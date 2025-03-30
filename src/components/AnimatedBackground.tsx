import React from 'react';

const AnimatedBackground = () => {
  return (
    <div className="fixed inset-0 overflow-hidden pointer-events-none z-0">
      <div className="cloud w-40 h-20 left-[10%] top-[15%] opacity-40 animate-float"></div>
      <div className="cloud w-56 h-24 right-[20%] top-[10%] opacity-30 animate-float" style={{ animationDelay: '1s' }}></div>
      <div className="cloud w-32 h-16 left-[30%] top-[25%] opacity-20 animate-float" style={{ animationDelay: '2s' }}></div>
      <div className="cloud w-48 h-20 right-[15%] top-[30%] opacity-25 animate-float" style={{ animationDelay: '3s' }}></div>
      <div className="cloud w-36 h-16 left-[15%] top-[40%] opacity-35 animate-float" style={{ animationDelay: '4s' }}></div>
      <div className="cloud w-44 h-18 right-[25%] top-[50%] opacity-30 animate-float" style={{ animationDelay: '5s' }}></div>
      <div className="cloud w-50 h-22 left-[20%] top-[60%] opacity-20 animate-float" style={{ animationDelay: '6s' }}></div>
      <div className="absolute inset-0 bg-gradient-to-b from-transparent via-background/30 to-background/90"></div>
    </div>
  );
};

export default AnimatedBackground;
