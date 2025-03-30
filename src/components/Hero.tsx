
import React from 'react';
import { Input } from '@/components/ui/input';
import { Search } from 'lucide-react';

interface HeroProps {
  searchQuery: string;
  onSearchChange: (query: string) => void;
}

const Hero = ({ searchQuery, onSearchChange }: HeroProps) => {
  return (
    <section className="relative px-6 py-16 md:py-20 flex flex-col items-center text-center z-10">
      <h1 className="text-4xl md:text-6xl mb-6 text-primary animate-fadeIn bg-clip-text text-transparent bg-gradient-to-r from-primary to-primary/80">
        Magical AI Prompts <br /> Inspired by Ghibli
      </h1>
      <p className="text-lg md:text-xl max-w-2xl text-muted-foreground mb-8 animate-fadeIn" style={{ animationDelay: '0.2s' }}>
        Explore a garden of whimsical, nostalgic, and nature-inspired AI art prompts that capture the enchanting essence of Studio Ghibli films.
      </p>
      
      <div className="w-full max-w-xl mx-auto mt-6 relative animate-fadeIn" style={{ animationDelay: '0.3s' }}>
        <div className="relative">
          <Search className="absolute left-3 top-3 h-5 w-5 text-muted-foreground/70" />
          <Input 
            placeholder="Search for magical prompts..." 
            className="pl-10 py-6 bg-muted/50 border-muted text-primary focus-visible:ring-ghibli-accent w-full rounded-full"
            value={searchQuery}
            onChange={(e) => onSearchChange(e.target.value)}
          />
        </div>
      </div>
    </section>
  );
};

export default Hero;
