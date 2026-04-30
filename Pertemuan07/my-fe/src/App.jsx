import { useState } from "react";
import Header from "./components/layout/Header";
import Sidebar from "./components/layout/Sidebar";
import Footer from "./components/layout/Footer";
import MahasiswaPage from "./pages/MahasiswaPages";
import DashboardPage from "./pages/DashboardPages";
import DataDiri from "./pages/DataDiri"; // Pastikan import ini sudah ada

export default function App() {
  const [activePage, setActivePage] = useState("dashboard");
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  // Fungsi untuk mendapatkan title berdasarkan activePage
  const getPageTitle = () => {
    switch (activePage) {
      case "dashboard": return "Dashboard";
      case "mahasiswa": return "Data Mahasiswa";
      case "data-diri": return "Data Diri";
      default: return "Dashboard";
    }
  };

  return (
    <div className="min-h-screen bg-slate-100 text-slate-800">
      <Sidebar
        activePage={activePage}
        isOpen={isSidebarOpen}
        onClose={() => setIsSidebarOpen(false)}
        onSelectPage={(pageId) => {
          setActivePage(pageId);
          setIsSidebarOpen(false);
        }}
      />

      <div className="flex min-h-screen flex-col md:pl-72">
        <Header
          pageTitle={getPageTitle()}
          onToggleSidebar={() => setIsSidebarOpen((prev) => !prev)}
        />

        <main className="flex-1 p-3 sm:p-4 md:p-6">
          <div className="w-full rounded-2xl border border-slate-200 bg-white p-4 shadow-sm md:p-6">
            {/* Conditional Rendering */}
            {activePage === "dashboard" && <DashboardPage />}
            {activePage === "mahasiswa" && <MahasiswaPage />}
            {activePage === "data-diri" && <DataDiri />}
          </div>
        </main>

        <Footer />
      </div>
    </div>
  );
}