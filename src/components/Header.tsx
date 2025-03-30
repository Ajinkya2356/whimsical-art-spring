
import React from 'react';
import { Moon, Sun, Search } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

const Header = () => {
  return (
    <header className="w-full py-4 px-6 md:px-10 flex flex-col md:flex-row justify-between items-center gap-4 relative z-10">
      <div className="flex items-center">
        <h1 className="text-3xl md:text-4xl text-ghibli-forest">Ghibli Prompt Garden</h1>
      </div>
      
      <div className="flex items-center gap-4 w-full md:w-auto">
        <div className="relative w-full md:w-64">
          <Input 
            placeholder="Search prompts..." 
            className="pl-9 rounded-full border-ghibli-forest/30 focus-visible:ring-ghibli-forest"
          />
          <Search className="absolute left-3 top-2.5 h-4 w-4 text-ghibli-forest/50" />
        </div>
        <Button variant="ghost" size="icon" className="rounded-full">
          <Sun className="h-5 w-5 rotate-0 scale-100 transition-all text-ghibli-accent" />
        </Button>
      </div>
    </header>
  );
};

export default Header;
