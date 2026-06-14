import { getUser } from "../services/auth";

export default function ProfilPage() {
  const user = getUser();

  return (
    <div className="space-y-5">
      <div className="mb-6">
        <h2 className="text-xl font-bold">Profil</h2>
        <p className="text-sm text-slate-500">Informasi akun yang sedang login.</p>
      </div>

      <div className="rounded-xl border border-slate-200 bg-white shadow-sm overflow-hidden">
        <div className="flex flex-col md:flex-row divide-y md:divide-y-0 md:divide-x divide-slate-100">
          <div className="flex-1 p-6">
            <h3 className="text-xs font-semibold uppercase text-slate-400 mb-1">
              USERNAME
            </h3>
            <p className="text-base font-medium text-slate-800">
              {user?.username || "-"}
            </p>
          </div>
          <div className="flex-1 p-6">
            <h3 className="text-xs font-semibold uppercase text-slate-400 mb-1">
              ROLE
            </h3>
            <p className="text-base font-medium text-slate-800">
              {user?.role || "-"}
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
