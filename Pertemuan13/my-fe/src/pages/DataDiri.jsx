// src/pages/DataDiri.jsx
import React from 'react';

// Komponen Ikon sederhana menggunakan teks (bisa diganti SVG jika perlu)
const Icon = ({ children }) => (
  <span className="flex items-center justify-center w-5 h-5 text-blue-500 mr-2 bg-blue-50 rounded-md">
    {children}
  </span>
);

const DetailCard = ({ icon, label, value }) => (
  <div className="flex items-center p-5 bg-white border border-slate-100 rounded-2xl shadow-sm hover:border-blue-100 hover:shadow-md transition-all duration-300">
    <Icon>{icon}</Icon>
    <div>
      <p className="text-xs font-semibold text-slate-400 uppercase tracking-wider">{label}</p>
      <p className="text-base font-semibold text-slate-900 mt-0.5">{value}</p>
    </div>
  </div>
);

export default function DataDiri() {
  return (
    <div className="space-y-8 max-w-7xl mx-auto p-1">
      {/* Header Halaman & Breadcrumb */}
      <div className="flex items-center justify-between">
        <h2 className="text-3xl font-extrabold text-slate-950 tracking-tight">Profil Mahasiswa</h2>
        <div className="text-sm text-slate-500 bg-slate-50 px-3 py-1.5 rounded-full border border-slate-100">
          Home / <span className="text-blue-600 font-medium">Data Diri</span>
        </div>
      </div>

      {/* --- SECTION 1: Banner & Profil Utama --- */}
      <div className="bg-white rounded-3xl border border-slate-100 shadow-xl shadow-slate-100/50 p-2">
        {/* Banner Area (Opsional, jika ingin menambahkan gambar latar belakang) */}
        <div className="h-32 bg-gradient-to-r from-blue-600 via-blue-500 to-indigo-600 rounded-2xl"></div>
        
        {/* Profil Content */}
        <div className="px-6 pb-6 -mt-16 flex flex-col md:flex-row items-center gap-6">
          {/* Avatar Ring */}
          <div className="w-32 h-32 bg-white rounded-full flex items-center justify-center border-8 border-white shadow-xl relative">
            <div className="w-full h-full bg-blue-600 rounded-full flex items-center justify-center text-white text-5xl font-black">
              {/* <img src={...} className="rounded-full object-cover"/> */}
              M
            </div>
            {/* Online Status Dot */}
            <span className="absolute bottom-2 right-2 w-6 h-6 bg-emerald-500 rounded-full border-4 border-white shadow-md"></span>
          </div>
          
          <div className="text-center md:text-left mt-4 md:mt-16 flex-1">
            <h3 className="text-2xl font-extrabold text-slate-950 tracking-tight">Muhammad Bagus Tri Atmaja</h3>
            <p className="text-lg font-medium text-blue-700 bg-blue-50 px-3 py-1 rounded-full inline-block mt-1">D4 Teknik Informatika</p>
            <p className="text-sm text-slate-500 mt-2">Aktif • Semester Ganjil 2026/2027</p>
          </div>
          
          <div className="mt-4 md:mt-16 flex gap-2">
            <button className="px-5 py-2.5 bg-blue-600 text-white rounded-xl font-semibold text-sm hover:bg-blue-700 transition shadow-sm">
              Edit Profil
            </button>
          </div>
        </div>
      </div>

      {/* --- SECTION 2: Grid Informasi Detail --- */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-5">
        <DetailCard icon="🆔" label="NPM" value="714240060" />
        <DetailCard icon="🏫" label="Kelas" value="D4 TI 2C" />
        <DetailCard icon="✉️" label="Email Politeknik" value="bagus@politeknik.ac.id" />
        <DetailCard icon="📅" label="Semester" value="Semester 4" />
        <DetailCard icon="🎓" label="Program Studi" value="D4 Teknik Informatika" />
        <DetailCard icon="💻" label="Mata Kuliah Utama" value="Web Service" />
      </div>

      {/* --- SECTION 3: Tentang & Catatan --- */}
      <div className="bg-white p-8 rounded-3xl border border-slate-100 shadow-lg shadow-slate-100/30">
        <div className="flex items-center mb-5">
          <Icon>📝</Icon>
          <h4 className="text-lg font-bold text-slate-950 uppercase tracking-wider">Tentang Halaman Ini</h4>
        </div>
        <div className="space-y-4 text-base text-slate-700 leading-relaxed max-w-4xl">
          <p>
            Halaman ini menampilkan informasi profil lengkap Anda sebagai mahasiswa. Ini adalah bagian dari Latihan Mandiri Pertemuan 7, 
            yang berfokus pada pengembangan antarmuka pengguna (UI) yang responsif dan terstruktur dengan React dan Tailwind CSS.
          </p>
          <p className="text-sm text-slate-500 bg-slate-50 p-4 rounded-xl border border-slate-100">
            <strong>Catatan Pengembangan:</strong> Tampilan ini dikembangkan untuk menjadi Final Project Anda, 
            pastikan semua data statis di sini akurat sebelum dikumpulkan. Anda juga dapat menghubungkan data ini ke API di masa depan.
          </p>
        </div>
      </div>
    </div>
  );
}