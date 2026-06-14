import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import Button from "../components/atoms/Button";
import TextInput from "../components/atoms/TextInput";
import FormField from "../components/molecules/FormField";
import { getToken, login, saveAuthSession } from "../services/auth";

export default function LoginPage() {
  const navigate = useNavigate();
  const [form, setForm] = useState({ username: "", password: "" });
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (getToken()) {
      navigate("/dashboard", { replace: true });
    }
  }, [navigate]);

  const handleChange = (event) => {
    const { name, value } = event.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      setLoading(true);
      const data = await login(form);
      saveAuthSession(data);
      await Swal.fire({
        title: "Berhasil",
        text: "Login berhasil.",
        icon: "success",
        confirmButtonText: "OK",
      });
      navigate("/dashboard", { replace: true });
    } catch (error) {
      await Swal.fire({
        title: "Gagal",
        text: error.response?.data?.message || "Username atau password salah.",
        icon: "error",
        confirmButtonText: "Tutup",
      });
    } finally {
      setLoading(false);
    }
  };

  return (
    <main className="flex min-h-screen items-center justify-center bg-slate-900 p-4">
      <div className="flex w-full max-w-5xl overflow-hidden rounded-3xl bg-white shadow-2xl">
        {/* Bagian Kiri (Informasi) */}
        <div className="hidden w-1/2 flex-col justify-center bg-gradient-to-br from-blue-600 to-indigo-900 p-12 text-white lg:flex">
          <p className="mb-2 text-xs font-bold uppercase tracking-widest text-blue-200">
            PERTEMUAN 12
          </p>
          <h1 className="mb-4 text-4xl font-bold">JWT Session & Profile</h1>
          <p className="mb-10 text-base leading-relaxed text-blue-100">
            Kelola sesi, profil, password, token, dan hak akses pengguna dalam satu aplikasi.
          </p>

          <div className="space-y-4">
            <div className="rounded-2xl border border-white/10 bg-white/10 p-5 backdrop-blur-sm">
              <p className="text-sm text-blue-50">
                Ubah password diamankan menggunakan bcrypt.
              </p>
            </div>
            <div className="rounded-2xl border border-white/10 bg-white/10 p-5 backdrop-blur-sm">
              <p className="text-sm text-blue-50">
                Respons 403 membedakan akun user dan admin.
              </p>
            </div>
          </div>
        </div>

        {/* Bagian Kanan (Form Login) */}
        <div className="w-full px-8 py-12 lg:w-1/2 lg:px-16 lg:py-16">
          <p className="mb-2 text-xs font-bold uppercase tracking-widest text-blue-600">
            LOGIN
          </p>
          <h2 className="mb-2 text-3xl font-bold text-slate-900">
            Masuk ke aplikasi
          </h2>
          <p className="mb-8 text-sm leading-relaxed text-slate-500">
            Gunakan akun admin untuk mengakses dan mengelola data mahasiswa.
          </p>

          <form onSubmit={handleSubmit} className="space-y-5">
            <FormField label="Username" htmlFor="username">
              <TextInput
                id="username"
                name="username"
                value={form.username}
                onChange={handleChange}
                placeholder="username"
                autoComplete="username"
                required
              />
            </FormField>

            <FormField label="Password" htmlFor="password">
              <TextInput
                id="password"
                name="password"
                type="password"
                value={form.password}
                onChange={handleChange}
                placeholder="password"
                autoComplete="current-password"
                required
              />
            </FormField>

            <Button type="submit" className="w-full py-3" disabled={loading}>
              {loading ? "Memproses..." : "Login"}
            </Button>
          </form>

          <p className="mt-8 text-center text-sm text-slate-600">
            Belum memiliki akun?{" "}
            <Link to="/register" className="font-semibold text-blue-600 hover:text-blue-700">
              Register
            </Link>
          </p>
        </div>
      </div>
    </main>
  );
}