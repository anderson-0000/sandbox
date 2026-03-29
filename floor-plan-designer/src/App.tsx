import Header from './components/ui/Header';
import Toolbar from './components/ui/Toolbar';
import Sidebar from './components/ui/Sidebar';
import ProjectStage from './components/canvas/Stage';

function App() {
  return (
    <div className="flex flex-col h-screen w-screen bg-transparent font-sans text-slate-900 overflow-hidden">
      <Header />
      <div className="flex flex-1 overflow-hidden relative">
        <Toolbar />
        <main className="flex-1 flex flex-col relative bg-white/40 shadow-[inset_0_2px_10px_rgba(0,0,0,0.02)]">
          <ProjectStage />
        </main>
        <Sidebar />
      </div>
    </div>
  );
}

export default App;
