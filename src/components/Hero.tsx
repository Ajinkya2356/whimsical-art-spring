
import React from 'react';
import { Button } from '@/components/ui/button';

const Hero = () => {
  return (
    <section className="relative px-6 py-16 md:py-24 flex flex-col items-center text-center z-10">
      <h1 className="text-4xl md:text-6xl mb-6 text-ghibli-forest animate-fadeIn">
        Magical AI Prompts <br /> Inspired by Ghibli
      </h1>
      <p className="text-lg md:text-xl max-w-2xl text-muted-foreground mb-8 animate-fadeIn" style={{ animationDelay: '0.2s' }}>
        Explore a garden of whimsical, nostalgic, and nature-inspired AI art prompts that capture the enchanting essence of Studio Ghibli films.
      </p>
      <div className="flex flex-col sm:flex-row gap-4 animate-fadeIn" style={{ animationDelay: '0.4s' }}>
        <Button className="ghibli-button">
          Explore Prompts
        </Button>
        <Button variant="outline" className="border-ghibli-forest text-ghibli-forest hover:bg-ghibli-forest/10">
          Learn More
        </Button>
      </div>
      
      <div className="mt-12 max-w-xs sm:max-w-sm md:max-w-md relative animate-fadeIn" style={{ animationDelay: '0.6s' }}>
        <div className="absolute -top-6 -left-6 w-24 h-24 bg-ghibli-magic/20 rounded-full blur-2xl"></div>
        <div className="absolute -bottom-8 -right-8 w-28 h-28 bg-ghibli-sunset/20 rounded-full blur-2xl"></div>
        <img 
          src="https://images.unsplash.com/photo-1506744038136-46273834b3fb" 
          alt="Ghibli-inspired landscape"
          className="rounded-2xl shadow-xl w-full object-cover"
        />
      </div>
    </section>
  );
};

export default Hero;
