
import React from 'react';

const Hero = () => {
  return (
    <section className="relative px-6 py-16 md:py-20 flex flex-col items-center text-center z-10">
      <h1 className="text-4xl md:text-6xl mb-10 text-primary animate-fadeIn bg-clip-text text-transparent bg-gradient-to-r from-primary to-primary/80 leading-tight py-7">
        Magical AI Prompts  
      </h1> 
      <p className="text-lg md:text-xl max-w-2xl text-muted-foreground mb-12 animate-fadeIn" style={{ animationDelay: '0.2s' }}>
        Explore a garden of whimsical, nostalgic, and nature-inspired AI art prompts that capture the enchanting essence of studio animated films.
      </p>
      
      
    </section>
  );
};

export default Hero;
