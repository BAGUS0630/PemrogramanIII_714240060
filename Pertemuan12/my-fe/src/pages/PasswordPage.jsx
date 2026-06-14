import { useState } from "react";
import Button from "../components/atoms/Button";
import TextInput from "../components/atoms/TextInput";
import Swal from "sweetalert2";

export default function PasswordPage() {
  const [oldPassword, setOldPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    if (newPassword !== confirmPassword) {
      Swal.fire({
        icon: "error",
        title: "Gagal",
        text: "Konfirmasi password baru tidak cocok!",
      });
      return;
    }
    // Simulate change password
    Swal.fire({
      icon: "success",
      title: "Berhasil",
      text: "Password berhasil diubah (Simulasi)",
    });
    setOldPassword("");
    setNewPassword("");
    setConfirmPassword("");
  };

  return (
    <div className="space-y-5">
      <div className="mb-6">
        <h2 className="text-xl font-bold">Ubah Password</h2>
        <p className="text-sm text-slate-500">
          Gunakan password lama untuk membuat password baru.
        </p>
      </div>

      <div className="rounded-xl border border-slate-200 bg-white p-6 shadow-sm">
        <form onSubmit={handleSubmit} className="space-y-4 max-w-md">
          <TextInput
            id="oldPassword"
            label="Password Lama"
            type="password"
            value={oldPassword}
            onChange={(e) => setOldPassword(e.target.value)}
            required
          />

          <TextInput
            id="newPassword"
            label="Password Baru"
            type="password"
            value={newPassword}
            onChange={(e) => setNewPassword(e.target.value)}
            required
          />

          <TextInput
            id="confirmPassword"
            label="Konfirmasi Password Baru"
            type="password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            required
          />

          <div className="pt-2">
            <Button type="submit" variant="primary">
              Ubah Password
            </Button>
          </div>
        </form>
      </div>
    </div>
  );
}
